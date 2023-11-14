package main

import "fmt"

/*
代理模式（Proxy Pattern）是一种结构型设计模式，它允许一个对象充当另一个对象的接口，以控制对这个对象的访问。
代理模式通常用于在访问一个对象时引入一些间接层，以实现控制、缓存、监控等功能。
*/

// ImageLoader 接口定义了图片加载的方法
type ImageLoader interface {
	DisplayImage() string
}

// RealImageLoader 实现了图片加载接口的具体类
type RealImageLoader struct {
	ImagePath string
}

func (r *RealImageLoader) DisplayImage() string {
	return fmt.Sprintf("Displaying image from path: %s", r.ImagePath)
}

// ImageLoaderProxy 代理类，实现了图片加载接口
type ImageLoaderProxy struct {
	RealImageLoader *RealImageLoader
}

func (p *ImageLoaderProxy) DisplayImage() string {
	// 在加载图片之前可以执行一些额外的操作，比如权限检查、缓存等
	return p.RealImageLoader.DisplayImage()
}

func main() {
	// 创建图片加载器代理
	imageLoaderProxy := &ImageLoaderProxy{
		RealImageLoader: &RealImageLoader{ImagePath: "/path/to/image.jpg"},
	}

	// 使用代理加载图片
	result := imageLoaderProxy.DisplayImage()
	fmt.Println(result)
	// Output: Displaying image from path: /path/to/image.jpg
}
