// +k8s:deepcopy-gen=package

// +groupName=samplecrd.k8s.io
package v1

// 其中，+k8s:deepcopy-gen=package 意思是，请为整个 v1 包里的所有类型定义自动生成 DeepCopy 方法；
// 而+groupName=samplecrd.k8s.io，则定义了这个包对应的 API 组的名字。
// 可以看到，这些定义在 doc.go 文件的注释，起到的是全局的代码生成控制的作用，所以也被称为 Global Tags。
