/*
 * Copyright 2017, gRPC Authors All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.grpc.services;

import com.google.common.annotations.VisibleForTesting;
import com.google.instrumentation.common.Duration;
import com.google.instrumentation.common.Timestamp;
import com.google.instrumentation.stats.DistributionAggregation;
import com.google.instrumentation.stats.DistributionAggregationDescriptor;
import com.google.instrumentation.stats.IntervalAggregation;
import com.google.instrumentation.stats.IntervalAggregation.Interval;
import com.google.instrumentation.stats.IntervalAggregationDescriptor;
import com.google.instrumentation.stats.MeasurementDescriptor;
import com.google.instrumentation.stats.MeasurementDescriptor.BasicUnit;
import com.google.instrumentation.stats.MeasurementDescriptor.MeasurementUnit;
import com.google.instrumentation.stats.Tag;
import com.google.instrumentation.stats.TagKey;
import com.google.instrumentation.stats.View;
import com.google.instrumentation.stats.View.DistributionView;
import com.google.instrumentation.stats.View.IntervalView;
import com.google.instrumentation.stats.ViewDescriptor;
import com.google.instrumentation.stats.ViewDescriptor.DistributionViewDescriptor;
import com.google.instrumentation.stats.ViewDescriptor.IntervalViewDescriptor;
import com.google.instrumentation.stats.proto.CensusProto;
import io.grpc.ExperimentalApi;
import io.grpc.instrumentation.v1alpha.StatsResponse;

/** Utility methods to support {@link MonitoringService}. */
@ExperimentalApi("https://github.com/grpc/grpc-java/issues/2776")
final class MonitoringUtil {

  private MonitoringUtil() {}

  /** Serialize a {@link View} and associated descriptors to a {@link StatsResponse}. */
  static StatsResponse buildCanonicalRpcStatsView(View view) {
    return StatsResponse.newBuilder()
        .setMeasurementDescriptor(
            serializeMeasurementDescriptor(view.getViewDescriptor().getMeasurementDescriptor()))
        .setViewDescriptor(serializeViewDescriptor(view.getViewDescriptor()))
        .setView(serializeView(view))
        .build();
  }

  @VisibleForTesting
  static CensusProto.MeasurementDescriptor serializeMeasurementDescriptor(
      MeasurementDescriptor measurementDescriptor) {
    return CensusProto.MeasurementDescriptor.newBuilder()
        .setName(measurementDescriptor.getName())
        .setDescription(measurementDescriptor.getDescription())
        .setUnit(serializeMeasurementUnit(measurementDescriptor.getUnit()))
        .build();
  }

  @VisibleForTesting
  static CensusProto.MeasurementDescriptor.MeasurementUnit serializeMeasurementUnit(
      MeasurementUnit unit) {
    CensusProto.MeasurementDescriptor.MeasurementUnit.Builder unitBuilder =
        CensusProto.MeasurementDescriptor.MeasurementUnit.newBuilder()
            .setPower10(unit.getPower10());
    for (BasicUnit basicUnit : unit.getNumerators()) {
      unitBuilder.addNumerators(serializeBasicUnit(basicUnit));
    }
    for (BasicUnit basicUnit : unit.getDenominators()) {
      unitBuilder.addDenominators(serializeBasicUnit(basicUnit));
    }
    return unitBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.MeasurementDescriptor.BasicUnit serializeBasicUnit(BasicUnit basicUnit) {
    switch (basicUnit) {
      case SCALAR:
        return CensusProto.MeasurementDescriptor.BasicUnit.SCALAR;
      case BITS:
        return CensusProto.MeasurementDescriptor.BasicUnit.BITS;
      case BYTES:
        return CensusProto.MeasurementDescriptor.BasicUnit.BYTES;
      case SECONDS:
        return CensusProto.MeasurementDescriptor.BasicUnit.SECONDS;
      case CORES:
        return CensusProto.MeasurementDescriptor.BasicUnit.CORES;
      default:
        return CensusProto.MeasurementDescriptor.BasicUnit.UNKNOWN;
    }
  }

  @VisibleForTesting
  static CensusProto.ViewDescriptor serializeViewDescriptor(ViewDescriptor viewDescriptor) {
    CensusProto.ViewDescriptor.Builder viewDescriptorBuilder =
        CensusProto.ViewDescriptor.newBuilder()
            .setName(viewDescriptor.getName())
            .setDescription(viewDescriptor.getDescription())
            .setMeasurementDescriptorName(viewDescriptor.getMeasurementDescriptor().getName());
    for (TagKey tagKey : viewDescriptor.getTagKeys()) {
      viewDescriptorBuilder.addTagKeys(tagKey.toString());
    }

    if (viewDescriptor instanceof DistributionViewDescriptor) {
      viewDescriptorBuilder.setDistributionAggregation(
          serializeDistributionAggregationDescriptor(
              ((DistributionViewDescriptor) viewDescriptor)
                  .getDistributionAggregationDescriptor()));
    } else {
      viewDescriptorBuilder.setIntervalAggregation(
          serializeIntervalAggregationDescriptor(
              ((IntervalViewDescriptor) viewDescriptor).getIntervalAggregationDescriptor()));
    }

    return viewDescriptorBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.DistributionAggregationDescriptor serializeDistributionAggregationDescriptor(
      DistributionAggregationDescriptor distributionAggregationDescriptor) {
    CensusProto.DistributionAggregationDescriptor.Builder distributionAggregationDescriptorBuilder =
        CensusProto.DistributionAggregationDescriptor.newBuilder();
    if (distributionAggregationDescriptor.getBucketBoundaries() != null) {
      distributionAggregationDescriptorBuilder.addAllBucketBounds(
          distributionAggregationDescriptor.getBucketBoundaries());
    }
    return distributionAggregationDescriptorBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.IntervalAggregationDescriptor serializeIntervalAggregationDescriptor(
      IntervalAggregationDescriptor intervalAggregationDescriptor) {
    CensusProto.IntervalAggregationDescriptor.Builder intervalAggregationDescriptorBuilder =
        CensusProto.IntervalAggregationDescriptor.newBuilder()
            .setNSubIntervals(intervalAggregationDescriptor.getNumSubIntervals());
    for (Duration intervalSize : intervalAggregationDescriptor.getIntervalSizes()) {
      intervalAggregationDescriptorBuilder.addIntervalSizes(serializeDuration(intervalSize));
    }
    return intervalAggregationDescriptorBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.Duration serializeDuration(Duration duration) {
    return CensusProto.Duration.newBuilder()
        .setSeconds(duration.getSeconds())
        .setNanos(duration.getNanos())
        .build();
  }

  @VisibleForTesting
  static CensusProto.View serializeView(View view) {
    CensusProto.View.Builder viewBuilder =
        CensusProto.View.newBuilder().setViewName(view.getViewDescriptor().getName());

    if (view instanceof DistributionView) {
      viewBuilder.setDistributionView(serializeDistributionView((DistributionView) view));
    } else {
      viewBuilder.setIntervalView(serializeIntervalView((IntervalView) view));
    }

    return viewBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.DistributionView serializeDistributionView(DistributionView distributionView) {
    CensusProto.DistributionView.Builder distributionViewBuilder =
        CensusProto.DistributionView.newBuilder();
    //TODO(ericgribkoff) Re-enable once getter methods are public in instrumentation.
    //distributionViewBuilder.setStart(serializeTimestamp(distributionView.getStart()))
    //distributionViewBuilder.setEnd(serializeTimestamp(distributionView.getEnd()));
    for (DistributionAggregation aggregation : distributionView.getDistributionAggregations()) {
      distributionViewBuilder.addAggregations(serializeDistributionAggregation(aggregation));
    }
    return distributionViewBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.Timestamp serializeTimestamp(Timestamp timestamp) {
    return CensusProto.Timestamp.newBuilder()
        .setSeconds(timestamp.getSeconds())
        .setNanos(timestamp.getNanos())
        .build();
  }

  @VisibleForTesting
  static CensusProto.DistributionAggregation serializeDistributionAggregation(
      DistributionAggregation aggregation) {
    CensusProto.DistributionAggregation.Builder aggregationBuilder =
        CensusProto.DistributionAggregation.newBuilder()
            .setCount(aggregation.getCount())
            .setMean(aggregation.getMean())
            .setSum(aggregation.getSum())
            .setRange(serializeRange(aggregation.getRange()));
    if (aggregation.getBucketCounts() != null) {
      aggregationBuilder.addAllBucketCounts(aggregation.getBucketCounts());
    }
    for (Tag tag : aggregation.getTags()) {
      aggregationBuilder.addTags(serializeTag(tag));
    }
    return aggregationBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.DistributionAggregation.Range serializeRange(
      DistributionAggregation.Range range) {
    CensusProto.DistributionAggregation.Range.Builder builder =
        CensusProto.DistributionAggregation.Range.newBuilder();
    if (range != null) {
      builder.setMin(range.getMin()).setMax(range.getMax());
    }
    return builder.build();
  }

  @VisibleForTesting
  static CensusProto.Tag serializeTag(Tag tag) {
    return CensusProto.Tag.newBuilder()
        .setKey(tag.getKey().toString())
        .setValue(tag.getValue().toString())
        .build();
  }

  @VisibleForTesting
  static CensusProto.IntervalView serializeIntervalView(IntervalView intervalView) {
    CensusProto.IntervalView.Builder intervalViewBuilder = CensusProto.IntervalView.newBuilder();
    for (IntervalAggregation aggregation : intervalView.getIntervalAggregations()) {
      intervalViewBuilder.addAggregations(serializeIntervalAggregation(aggregation));
    }
    return intervalViewBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.IntervalAggregation serializeIntervalAggregation(
      IntervalAggregation aggregation) {
    CensusProto.IntervalAggregation.Builder aggregationBuilder =
        CensusProto.IntervalAggregation.newBuilder();
    for (Interval interval : aggregation.getIntervals()) {
      aggregationBuilder.addIntervals(serializeInterval(interval));
    }
    for (Tag tag : aggregation.getTags()) {
      aggregationBuilder.addTags(serializeTag(tag));
    }
    return aggregationBuilder.build();
  }

  @VisibleForTesting
  static CensusProto.IntervalAggregation.Interval serializeInterval(Interval interval) {
    return CensusProto.IntervalAggregation.Interval.newBuilder()
        .setIntervalSize(serializeDuration(interval.getIntervalSize()))
        .setCount(interval.getCount())
        .setSum(interval.getSum())
        .build();
  }
}
