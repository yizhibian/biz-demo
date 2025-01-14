// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"

	"github.com/cloudwego/biz-demo/easy_note/cmd/note/dal/db"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demonote"
)

type CreateNoteService struct {
	ctx context.Context
}

// NewCreateNoteService new CreateNoteService
func NewCreateNoteService(ctx context.Context) *CreateNoteService {
	return &CreateNoteService{ctx: ctx}
}

// CreateNote create note info
func (s *CreateNoteService) CreateNote(req *demonote.CreateNoteRequest) error {

	klog.Info("here we gone make the wrong")
	noteModel := &db.Note{
		//插入重复数据以制造回滚
		Model: gorm.Model{
			ID: 1,
		},
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}
	return db.CreateNote(s.ctx, []*db.Note{noteModel})
}
