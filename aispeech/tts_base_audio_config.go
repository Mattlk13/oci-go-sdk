// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TtsBaseAudioConfig Use this audio config for saving the audio response at specified path.
type TtsBaseAudioConfig struct {

	// Specify the path where you want to save the audio response.
	SavePath *string `mandatory:"true" json:"savePath"`
}

func (m TtsBaseAudioConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TtsBaseAudioConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TtsBaseAudioConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTtsBaseAudioConfig TtsBaseAudioConfig
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeTtsBaseAudioConfig
	}{
		"BASE_AUDIO_CONFIG",
		(MarshalTypeTtsBaseAudioConfig)(m),
	}

	return json.Marshal(&s)
}
