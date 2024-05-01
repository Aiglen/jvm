package classfile

/*
	constant_class_info {
		tag		u1
		index 	u2
	}
*/
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) String() string {
	return self.cp.getUtf8(self.nameIndex)
}
