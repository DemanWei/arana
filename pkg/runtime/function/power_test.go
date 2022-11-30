/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package function

import (
	"context"
	"fmt"
	"testing"
)

import (
	gxbig "github.com/dubbogo/gost/math/big"

	"github.com/stretchr/testify/assert"
)

import (
	"github.com/arana-db/arana/pkg/proto"
)

func TestPower(t *testing.T) {
	fn := proto.MustGetFunc(FuncPower)
	assert.Equal(t, 2, fn.NumInput())

	type tt struct {
		in  proto.Value
		in2 proto.Value
		out string
	}

	mustDecimal := func(s string) *gxbig.Decimal {
		d, _ := gxbig.NewDecFromString(s)
		return d
	}

	for _, it := range []tt{
		{int8(12), 0, "1"},
		{12.34, 2, "152.2756"},
		{float64(-1.999), 2, "3.9960010000000006"},
		{mustDecimal("-5"), -2, "0.04"},
		{"foobar", -1, "NaN"},
		{"1", "foobar", "NaN"},
	} {
		t.Run(it.out, func(t *testing.T) {
			out, err := fn.Apply(context.Background(), proto.ToValuer(it.in), proto.ToValuer(it.in2))
			assert.NoError(t, err)
			assert.Equal(t, it.out, fmt.Sprint(out))
		})
	}
}