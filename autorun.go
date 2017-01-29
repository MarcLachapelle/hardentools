/*
   Hardentools
   Copyright (C) 2017  Claudio Guarnieri, Mariano Graziano

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
    "golang.org/x/sys/windows/registry"
)

/*
 * Disables Windows AutoRun, still needs thourough testing!
 *
 *  // HKCU\Software\Microsoft\Windows\CurrentVersion\Policies\Explorer!NoDriveTypeAutoRun
 *  // HKCU\Software\Microsoft\Windows\CurrentVersion\Policies\Explorer!NoAutorun
 *  // HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\AutoplayHandlers!DisableAutoplay 1
 */
func trigger_autorun(enable bool) {
    key, _, _ := registry.CreateKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Policies\\Explorer", registry.WRITE)
    key2, _, _ := registry.CreateKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\AutoplayHandlers", registry.WRITE)

    if enable {
        events.AppendText("* Enabling AutoRun and AutoPlay\n")
        key.DeleteValue("NoDriveTypeAutoRun")
        key.DeleteValue("NoAutorun")
        key2.DeleteValue("DisableAutoplay")
    } else {
        events.AppendText("* Disabling AutoRun and AutoPlay\n")
        key.SetDWordValue("NoDriveTypeAutoRun", 0xb5)
        key.SetDWordValue("NoAutorun", 1)
        key2.SetDWordValue("DisableAutoplay", 1)
    }

    key.Close()
    key2.Close()
}
