# Randpic

一个简单的命令行工具，用于从 Unsplash 获取随机图片链接。

## 功能特点

- 获取 Unsplash 随机图片链接
- 支持按主题搜索图片
- 支持复制结果到剪贴板
- 简单易用的命令行界面

## 使用前准备

1. 注册 [Unsplash Developers](https://unsplash.com/developers) 账号
2. 创建应用并获取 Access Key
3. 设置环境变量：
   ```bash
   export UNSPLASH_ACCESS_KEY="your-access-key"
   ```

## 使用方法

### 基本使用

获取一张随机图片：
```bash
randpic
```

### 按主题搜索

指定主题获取随机图片：
```bash
randpic -q "nature"    # 获取自然主题的随机图片
randpic --query "cat"  # 获取猫咪主题的随机图片
```

### 显示详细信息

显示图片的所有尺寸链接：
```bash
randpic -d            # 显示随机图片的所有尺寸链接
randpic --detail     # 同上
```

### 复制到剪贴板

将结果直接复制到剪贴板：
```bash
randpic -c            # 复制随机图片链接到剪贴板
randpic --clipboard   # 同上
```

### 组合使用

```bash
randpic -q "nature" -d  # 获取自然主题的随机图片，并显示所有尺寸链接
randpic -q "cat" -c     # 获取猫咪主题的随机图片，并复制到剪贴板
randpic -d -c           # 显示所有尺寸链接，并复制到剪贴板
```

## 安装

```bash
go install github.com/weirwei/randpic@latest
```

## 贡献指南

欢迎为 Randpic 做出贡献！以下是一些贡献的方式：

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 问题反馈

如果你在使用过程中遇到任何问题或有新的功能建议，欢迎：

- [提交 Issue](https://github.com/weirwei/randpic/issues)
- [提交 Pull Request](https://github.com/weirwei/randpic/pulls)

## 更新日志

### v1.1.0 - 2025-03-14

- 新增复制到剪贴板功能
- 支持使用 -c/--clipboard 标志

### v1.0.0 - 2025-03-12

- 初始版本发布
- 支持获取随机图片
- 支持按主题搜索
- 支持显示详细信息

## 致谢

- [Unsplash](https://unsplash.com/) - 提供优质的图片资源

## 作者

- [@weirwei](https://github.com/weirwei)