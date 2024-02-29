package proto

// Test_Path_Not_Found not.found
var Test_Path_Not_Found = NewPath(PathType_Library, []string{"__not__", "__found__"})

// Test_Path_Rest_From rest.from
var Test_Path_Rest_From = NewPath(PathType_Library, []string{"rest", "from"})

// Test_Path_Rest_Method rest.method
var Test_Path_Rest_Method = NewPath(PathType_Library, []string{"rest", "method"})

// Test_Path_Rest_Handler rest.handler
var Test_Path_Rest_Handler = NewPath(PathType_Library, []string{"rest", "handler"})

// Test_Path_Rest_Key rest.key
var Test_Path_Rest_Key = NewPath(PathType_Library, []string{"rest", "key"})

// Test_Path_Rest_Key_ rest.key_
var Test_Path_Rest_Key_ = NewPath(PathType_Library, []string{"rest", "key_"})

// Test_Path_App_Key21 app.key2.key21
var Test_Path_App_Key21 = NewPath(PathType_Application, []string{"key2", "key21"})

// Test_Path_App_Message app.message
var Test_Path_App_Message = NewPath(PathType_Application, []string{"__message__"})

// Test_Path_Lib1_Key11 lib1.key1.key11
var Test_Path_Lib1_Key11 = NewPath(PathType_Library, []string{"lib1", "key1", "key11"})

// Test_AttributeConfigure_App_Key21 tag["app.key21"] = val(app.key2.key21)
var Test_AttributeConfigure_App_Key21 = NewAttributeConfigure("app.key21", Test_Path_App_Key21)

// Test_AttributeConfigure_App_Message tag["app.message"] = val(app.message)
var Test_AttributeConfigure_App_Message = NewAttributeConfigure("app.message", Test_Path_App_Message)

// Test_AttributeConfigure_Lib1_Key11 tag["lib1.key11"] = val(lib1.key1.key11)
var Test_AttributeConfigure_Lib1_Key11 = NewAttributeConfigure("lib1.key11", Test_Path_Lib1_Key11)

// Test_AttributeConfigure_Rest_From tag["rest.from"] = val(rest.from)
var Test_AttributeConfigure_Rest_From = NewAttributeConfigure("rest.from", Test_Path_Rest_From)

// Test_AttributeConfigure_Rest_Method tag["rest.method"] = val(rest.method)
var Test_AttributeConfigure_Rest_Method = NewAttributeConfigure("method", Test_Path_Rest_Method)

// Test_AttributeConfigure_Rest_Handler tag["rest.handler"] = val(rest.handler)
var Test_AttributeConfigure_Rest_Handler = NewAttributeConfigure("handler", Test_Path_Rest_Handler)

// Test_AttributeConfigure_Rest_Key tag["rest.key"] = val(rest.key)
var Test_AttributeConfigure_Rest_Key = NewAttributeConfigure("rest.key", Test_Path_Rest_Key)

// Test_AttributeConfigure_Rest_Key_ tag["rest.key_"] = val(rest.key_)
var Test_AttributeConfigure_Rest_Key_ = NewAttributeConfigure("rest.key_", Test_Path_Rest_Key_)
