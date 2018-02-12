package classfile

/*
cp_info{
u1 tag //用来标识当前常量类型
u1 info[]
}
因为常量池中表项的具体类型不确定，因此定义一个接口，根据实际情况读取字节流
*/
// 定义常量池的项的数据类型；即tag的枚举值
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Dubble             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

/*
读取常量池的项
 */
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// 读取常量池类型
	//tag := reader.readUint8()
	// TODO 根据类型生成常量结构体
	// TODO 读取具体的常量信息
	return nil
}

/*
根据常量的tag生成对应类型的常量结构体
*/
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	return nil
}