package v1

//types.go 文件。它的作用就是定义一个 Network 类型到底有哪些字段
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
其中，+genclient 的意思是：请为下面这个 API 资源类型生成对应的 Client 代码（这个 Client，我马上会讲到）。
而 +genclient:noStatus 的意思是：这个 API 资源类型定义里，没有 Status 字段。
否则，生成的 Client 就会自动带上 UpdateStatus 方法。
如果你的类型定义包括了 Status 字段的话，就不需要这句 +genclient:noStatus 注释了

在 Global Tags 里已经定义了为所有类型生成 DeepCopy 方法，
所以这里就不需要再显式地加上 +k8s:deepcopy-gen=true 了。
当然，这也就意味着你可以用 +k8s:deepcopy-gen=false 来阻止为某些类型生成 DeepCopy。

+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object的注释。
它的意思是，请在生成 DeepCopy 的时候，实现 Kubernetes 提供的 runtime.Object 接口。
否则，在某些版本的 Kubernetes 里，你的这个类型定义会出现编译错误。
这是一个固定的操作，记住即可。
*/

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Network describes a Network resource
type Network struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec NetworkSpec `json:"spec"`
}

/* 你可以看到 Network 类型定义方法跟标准的 Kubernetes 对象一样，
都包括了 TypeMeta（API 元数据）和 ObjectMeta（对象元数据）字段。
而其中的 Spec 字段，就是需要我们自己定义的部分。
所以，在 networkspec 里，我定义了 Cidr 和 Gateway 两个字段。
其中，每个字段最后面的部分比如json:"cidr"，指的就是这个字段被转换成 JSON 格式之后的名字，
也就是 YAML 文件里的字段名字。
*/

// NetworkSpec is the spec for a Network resource
type NetworkSpec struct {
	// Cidr and Gateway are example custom spec fields
	//
	// this is where you would put your custom resource data
	Cidr    string `json:"cidr"`
	Gateway string `json:"gateway"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList is a list of Network resources
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Network `json:"items"`
}

/*
你还需要定义一个 NetworkList 类型，用来描述一组 Network 对象应该包括哪些字段。
之所以需要这样一个类型，是因为在 Kubernetes 中，获取所有 X 对象的 List() 方法，返回值都是List 类型，
而不是 X 类型的数组。
*/
