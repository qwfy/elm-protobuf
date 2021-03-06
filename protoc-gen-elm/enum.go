package main

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func (fg *FileGenerator) GenerateEnumDefinition(prefix string, inEnum *descriptor.EnumDescriptorProto) error {
	typeName := prefix + inEnum.GetName()
	fg.P("")
	fg.P("")
	fg.P("type %s", typeName)
	{
		fg.In()
		leading := "="
		for _, enumValue := range inEnum.GetValue() {
			// TODO: Convert names to CamelCase.
			fg.P("%s %s -- %d", leading, prefix+elmEnumValueName(enumValue.GetName()), enumValue.GetNumber())
			leading = "|"
		}
		fg.Out()
	}

    fg.P("")
	fg.P("")
	fg.P("intOf%s : %s -> Int", typeName, typeName)
	fg.P("intOf%s x =", typeName)
	{
	    fg.In()
        fg.P("case x of")
        fg.In()
		for _, enumValue := range inEnum.GetValue() {
			// TODO: Convert names to CamelCase.
			fg.P("%s -> %d", elmEnumValueName(enumValue.GetName()), enumValue.GetNumber())
		}
		fg.Out()
		fg.Out()
	}
	return nil
}

func (fg *FileGenerator) GenerateEnumDecoder(prefix string, inEnum *descriptor.EnumDescriptorProto) error {
	typeName := prefix + inEnum.GetName()
	decoderName := decoderName(typeName)
	fg.P("")
	fg.P("")
    fg.P("decode%s : JD.Decoder %s", typeName, typeName)
    fg.P("decode%s = %s", typeName, decoderName)
    fg.P("")
	fg.P("")
	fg.P("%s : JD.Decoder %s", decoderName, typeName)
	fg.P("%s =", decoderName)
	{
		fg.In()
		fg.P("let")
		{
			fg.In()
			fg.P("lookup s =")
			fg.In()
			fg.P("case s of")
			{
				fg.In()
				for _, enumValue := range inEnum.GetValue() {
					fg.P("%q ->", enumValue.GetName())
					fg.In()
					fg.P("%s", prefix+elmEnumValueName(enumValue.GetName()))
					fg.P("")
					fg.Out()
				}
				// TODO: This should fail instead.
				fg.P("_ ->")
				fg.In()
				fg.P("%s", prefix+elmEnumValueName(inEnum.GetValue()[0].GetName()))
				fg.Out()
				fg.Out()
			}
			fg.Out()
			fg.Out()
		}
		fg.P("in")
		{
			fg.In()
			fg.P("JD.map lookup JD.string")
			fg.Out()
		}
		fg.Out()
	}

	defaultName := defaultEnumValue(typeName)
	fg.P("")
	fg.P("")
	fg.P("%s : %s", defaultName, typeName)
	fg.P("%s = %s", defaultName, prefix+elmEnumValueName(inEnum.GetValue()[0].GetName()))
	return nil
}

func (fg *FileGenerator) GenerateEnumEncoder(prefix string, inEnum *descriptor.EnumDescriptorProto) error {
	typeName := prefix + inEnum.GetName()
	argName := "v"
	fg.P("")
	fg.P("")
    fg.P("encode%s : %s -> JE.Value", typeName, typeName)
    fg.P("encode%s = %s", typeName, encoderName(typeName))
    fg.P("")
	fg.P("")
	fg.P("%s : %s -> JE.Value", encoderName(typeName), typeName)
	fg.P("%s %s =", encoderName(typeName), argName)
	{
		fg.In()
		fg.P("let")
		{
			fg.In()
			fg.P("lookup s =")
			fg.In()
			fg.P("case s of")
			{
				fg.In()
				for _, enumValue := range inEnum.GetValue() {
					fg.P("%s ->", prefix+elmEnumValueName(enumValue.GetName()))
					fg.In()
					fg.P("%q", enumValue.GetName())
					fg.P("")
					fg.Out()
				}
				fg.Out()
			}
			fg.Out()
			fg.Out()
		}
		fg.P("in")
		{
			fg.In()
			fg.P("JE.string <| lookup %s", argName)
			fg.Out()
		}
		fg.Out()
	}
	return nil
}
