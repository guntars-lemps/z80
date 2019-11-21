/*

Copyright (c) 2010 Andrea Fazzi

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*/

package z80

func (z80 *Z80) BC() uint16 {
	return z80.bc.get()
}

func (z80 *Z80) SetBC(value uint16) {
	z80.bc.set(value)
}

func (z80 *Z80) DecBC() {
	z80.bc.dec()
}

func (z80 *Z80) IncBC() {
	z80.bc.inc()
}

func (z80 *Z80) DE() uint16 {
	return z80.de.get()
}

func (z80 *Z80) SetDE(value uint16) {
	z80.de.set(value)
}

func (z80 *Z80) DecDE() {
	z80.de.dec()
}

func (z80 *Z80) IncDE() {
	z80.de.inc()
}

func (z80 *Z80) HL() uint16 {
	return z80.hl.get()
}

func (z80 *Z80) SetHL(value uint16) {
	z80.hl.set(value)
}

func (z80 *Z80) DecHL() {
	z80.hl.dec()
}

func (z80 *Z80) IncHL() {
	z80.hl.inc()
}

func (z80 *Z80) BC_() uint16 {
	return z80.bc_.get()
}

func (z80 *Z80) SetBC_(value uint16) {
	z80.bc_.set(value)
}

func (z80 *Z80) DecBC_() {
	z80.bc_.dec()
}

func (z80 *Z80) IncBC_() {
	z80.bc_.inc()
}

func (z80 *Z80) DE_() uint16 {
	return z80.de_.get()
}

func (z80 *Z80) SetDE_(value uint16) {
	z80.de_.set(value)
}

func (z80 *Z80) DecDE_() {
	z80.de_.dec()
}

func (z80 *Z80) IncDE_() {
	z80.de_.inc()
}

func (z80 *Z80) HL_() uint16 {
	return z80.hl_.get()
}

func (z80 *Z80) SetHL_(value uint16) {
	z80.hl_.set(value)
}

func (z80 *Z80) DecHL_() {
	z80.hl_.dec()
}

func (z80 *Z80) IncHL_() {
	z80.hl_.inc()
}

func (z80 *Z80) IX() uint16 {
	return z80.ix.get()
}

func (z80 *Z80) SetIX(value uint16) {
	z80.ix.set(value)
}

func (z80 *Z80) DecIX() {
	z80.ix.dec()
}

func (z80 *Z80) IncIX() {
	z80.ix.inc()
}

func (z80 *Z80) IY() uint16 {
	return z80.iy.get()
}

func (z80 *Z80) SetIY(value uint16) {
	z80.iy.set(value)
}

func (z80 *Z80) DecIY() {
	z80.iy.dec()
}

func (z80 *Z80) IncIY() {
	z80.iy.inc()
}
