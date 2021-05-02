// Copyright (C) 2008-2019 by Nicolas Piganeau and the TS2 TEAM
// (See AUTHORS file)
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the
// Free Software Foundation, Inc.,
// 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.

package lines

import "ts2-server/simulation"

// StandardManager is a lines manager that never fails.
type StandardManager struct{}

// IsFailed returns whether the track circuit of the given line item is failed or not
func (sm StandardManager) IsFailed(p *simulation.LineItem) bool {
	return false
}

// Name returns a description of this manager that is used for the UI.
func (sm StandardManager) Name() string {
	return "Standard Manager"
}

var _ simulation.LineItemManager = StandardManager{}

func init() {
	simulation.RegisterLineItemManager(StandardManager{})
}