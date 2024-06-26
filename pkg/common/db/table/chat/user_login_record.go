// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chat

import (
	"context"
	"time"
)

type UserLoginRecord struct {
	UserID    string    `bson:"user_id"`
	LoginTime time.Time `bson:"login_time"`
	IP        string    `bson:"ip"`
	DeviceID  string    `bson:"device_id"`
	Platform  string    `bson:"platform"`
}

func (UserLoginRecord) TableName() string {
	return "user_login_records"
}

type UserLoginRecordInterface interface {
	Create(ctx context.Context, records ...*UserLoginRecord) error
	CountTotal(ctx context.Context, before *time.Time) (int64, error)
	CountRangeEverydayTotal(ctx context.Context, start *time.Time, end *time.Time) (map[string]int64, int64, error)
}
