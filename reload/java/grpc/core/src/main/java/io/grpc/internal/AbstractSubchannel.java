/*
 * Copyright 2016, gRPC Authors All rights reserved.
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

package io.grpc.internal;

import io.grpc.LoadBalancer;
import io.grpc.internal.Channelz.ChannelStats;
import javax.annotation.Nullable;

/**
 * The base interface of the Subchannels returned by {@link
 * io.grpc.LoadBalancer.Helper#createSubchannel}.
 */
abstract class AbstractSubchannel extends LoadBalancer.Subchannel
    implements Instrumented<ChannelStats> {
  private final LogId logId = LogId.allocate(getClass().getName());

  /**
   * Same as {@link InternalSubchannel#obtainActiveTransport}.
   */
  @Nullable
  abstract ClientTransport obtainActiveTransport();

  @Override
  public LogId getLogId() {
    return logId;
  }
}
