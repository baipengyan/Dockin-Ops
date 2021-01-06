/*
 * Copyright (C) @2021 Webank Group Holding Limited
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package ctrl

import (
	"github.com/webankfintech/dockin-opserver/internal/cache/keys"
	"github.com/webankfintech/dockin-opserver/internal/cache/redis"
	"github.com/webankfintech/dockin-opserver/internal/log"
)

type Version struct {
	redisClient *redis.RedisClient
}

func (v *Version) LoadVersionFromCache() string {
	version, err := v.redisClient.Get(keys.GetVersionKey())
	if err != nil {
		log.Logger.Warnf("redis get version from key:%s failed,err=%s", keys.GetVersionKey(), err.Error())
		return ""
	}

	return version.(string)
}

func (v *Version) UpdateVersion(version interface{}) error {
	return v.redisClient.Set(keys.GetVersionKey(), version, 0)
}