package reaction

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// A list of type definitions from proto messages, in the order of their definition.

// Conditions

type ConditionMessage cb.ConditionMessage
type ConditionLogic cb.ConditionLogic
type ConditionNode cb.ConditionNode

type PrerequisiteMessage cb.PrerequisiteMessage
type PrerequisiteLogic cb.PrerequisiteLogic
type PrerequisiteNode cb.PrerequisiteNode
type PrerequisiteTree struct {
	cb.PrerequisiteTree
	Index map[string]*PrerequisiteNode // <event name, PrerequisiteNode>
}

// Snapshot

type PrerequisiteSnapshot cb.PrerequisiteSnapshot

// Params for reaction operators

type FaultDelayParam cb.FaultDelayParam
type TrafficBalanceParam cb.TrafficBalanceParam
type TrafficRoutingParam cb.TrafficRoutingParam

type Configure struct {
	Name    string
	Type    cb.ReactionType
	Params  interface{} // isReactionConfigure_Params
	PreTree *PrerequisiteTree
}

type ReactionConfigure_FaultDelay cb.ReactionConfigure_FaultDelay
type ReactionConfigure_TrafficBalance cb.ReactionConfigure_TrafficBalance
type ReactionConfigure_TrafficRouting cb.ReactionConfigure_TrafficRouting
