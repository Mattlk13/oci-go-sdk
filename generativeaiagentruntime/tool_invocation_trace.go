// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Client API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents Client API to create and manage client chat sessions. A session represents an interactive conversation initiated by a user through an API to engage with an agent. It involves a series of exchanges where the user sends queries or prompts, and the agent responds with relevant information, actions, or assistance based on the user's input. The session persists for the duration of the interaction, maintaining context and continuity to provide coherent and meaningful responses throughout the conversation.
// For creating and managing agents, knowledge bases, data sources, endpoints, and data ingestion jobs see the /EN/generative-ai-agents/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagentruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ToolInvocationTrace The trace information about the tool selection from multiple tools.
type ToolInvocationTrace struct {

	// Unique identifier for the event (UUID).
	Key *string `mandatory:"false" json:"key"`

	// Identifier of the parent event, if applicable (UUID).
	ParentKey *string `mandatory:"false" json:"parentKey"`

	Source *SourceDetails `mandatory:"false" json:"source"`

	// The date and time that the trace was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Timestamp for when the event ended (In RFC 3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The ID of the selected tool based on the user query.
	ToolId *string `mandatory:"false" json:"toolId"`

	// The display name of the selected tool.
	ToolName *string `mandatory:"false" json:"toolName"`

	// The invocation details related to the selected tool.
	InvocationDetails *string `mandatory:"false" json:"invocationDetails"`
}

// GetKey returns Key
func (m ToolInvocationTrace) GetKey() *string {
	return m.Key
}

// GetParentKey returns ParentKey
func (m ToolInvocationTrace) GetParentKey() *string {
	return m.ParentKey
}

// GetSource returns Source
func (m ToolInvocationTrace) GetSource() *SourceDetails {
	return m.Source
}

// GetTimeCreated returns TimeCreated
func (m ToolInvocationTrace) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeFinished returns TimeFinished
func (m ToolInvocationTrace) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

func (m ToolInvocationTrace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ToolInvocationTrace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ToolInvocationTrace) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeToolInvocationTrace ToolInvocationTrace
	s := struct {
		DiscriminatorParam string `json:"traceType"`
		MarshalTypeToolInvocationTrace
	}{
		"TOOL_INVOCATION_TRACE",
		(MarshalTypeToolInvocationTrace)(m),
	}

	return json.Marshal(&s)
}
