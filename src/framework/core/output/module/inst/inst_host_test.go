/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package inst_test

import (
	"configcenter/src/framework/common"
	"configcenter/src/framework/core/config"
	"configcenter/src/framework/core/output/module/client"
	"configcenter/src/framework/core/output/module/inst"
	"configcenter/src/framework/core/output/module/model"
	//"configcenter/src/framework/core/types"
	"testing"
)

func TestHostManager(t *testing.T) {

	client.NewForConfig(config.Config{"supplierAccount": "0", "user": "build_user", "ccaddress": "http://test.apiserver:8080"}, nil)

	clsItem, err := model.FindClassificationsByCondition(common.CreateCondition().Field("bk_classification_id").Eq("bk_host_manage"))
	if nil != err {
		t.Errorf("failed to find classifications, %s", err.Error())
		return
	}

	if nil == clsItem {
		t.Errorf("not found the host classification")
		return
	}

	clsItem.ForEach(func(item model.Classification) {

		modelIter, err := item.FindModelsByCondition(common.CreateCondition().Field("bk_obj_id").Eq("host"))
		if nil != err {
			t.Errorf("failed to search classification, %s", err.Error())
			return
		}

		if nil == modelIter {
			t.Log("not found the model")
			return
		}

		// deal host model
		modelIter.ForEach(func(modelItem model.Model) {

			// create host
			hostInst, err := inst.CreateInst(modelItem)
			if nil != err {
				t.Errorf("failed to create host ")
				return
			}

			// Only test
			t.Logf("model name:%s", hostInst.GetModel().GetName())

			// set host value
			err = hostInst.SetValue("test", "test")
			if nil != err {
				t.Errorf("failed to set value, %s", err.Error())
				return
			}

			// save host info
			err = hostInst.Save()

			if nil != err {
				t.Errorf("failed to save ,%s", err.Error())
			}

		})

	})

}
