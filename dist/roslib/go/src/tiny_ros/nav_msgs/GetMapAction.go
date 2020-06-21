package nav_msgs

import (
    "nav_msgs/GetMapActionGoal"
    "nav_msgs/GetMapActionResult"
    "nav_msgs/GetMapActionFeedback"
)

type GetMapAction struct {
    Go_action_goal nav_msgs.GetMapActionGoal `json:"action_goal"`
    Go_action_result nav_msgs.GetMapActionResult `json:"action_result"`
    Go_action_feedback nav_msgs.GetMapActionFeedback `json:"action_feedback"`
}

func NewGetMapAction() (*GetMapAction) {
    newGetMapAction := new(GetMapAction)
    newGetMapAction.Go_action_goal = nav_msgs.NewGetMapActionGoal()
    newGetMapAction.Go_action_result = nav_msgs.NewGetMapActionResult()
    newGetMapAction.Go_action_feedback = nav_msgs.NewGetMapActionFeedback()
    return newGetMapAction
}

func (self *GetMapAction) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_action_goal.Go_serialize(buff[offset:])
    offset += self.Go_action_result.Go_serialize(buff[offset:])
    offset += self.Go_action_feedback.Go_serialize(buff[offset:])
    return offset
}

func (self *GetMapAction) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_action_goal.Go_deserialize(buff[offset:])
    offset += self.Go_action_result.Go_deserialize(buff[offset:])
    offset += self.Go_action_feedback.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetMapAction) Go_serializedLength() (int) {
    length := 0
    length += self.Go_action_goal.Go_serializedLength()
    length += self.Go_action_result.Go_serializedLength()
    length += self.Go_action_feedback.Go_serializedLength()
    return length
}

func (self *GetMapAction) Go_echo() (string) { return "" }
func (self *GetMapAction) Go_getType() (string) { return "nav_msgs/GetMapAction" }
func (self *GetMapAction) Go_getMD5() (string) { return "10a4e277d7b8e53bfc3df54d98b3edb1" }
func (self *GetMapAction) Go_getID() (uint32) { return 0 }
func (self *GetMapAction) Go_setID(id uint32) { }
