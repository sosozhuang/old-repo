/*
Copyright 2014 Huawei Technologies Co., Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package module

import (
	"fmt"

	"github.com/Huawei/containerops/pilotage/models"

	log "github.com/Sirupsen/logrus"
)

func RecordOutcom(workflowID, fromWorkflowID, stageID, fromStageID, actionID, fromActionID, sequence, evnetId int64, status bool, result, output string) error {
	outcome := new(models.Outcome)
	outcome.Workflow = workflowID
	outcome.Sequence = sequence
	outcome.Stage = stageID
	outcome.Action = actionID
	outcome.RealWorkflow = fromWorkflowID
	outcome.RealStage = fromStageID
	outcome.RealAction = fromActionID
	outcome.Event = evnetId
	outcome.Status = status
	outcome.Result = result
	outcome.Output = output

	err := outcome.Save()
	if err != nil {
		log.Error("[outcome's RecordOutcome]:error when save outcome info:", fmt.Sprintf("%#v", outcome), " ===>error is:", err.Error())
		return err
	}

	return nil
}
