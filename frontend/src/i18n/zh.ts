const ZH = {
    sideBar: {
        tabsTitle: ['建议', '文章', '设置']
    }, inquiry: {
        header: {
            button: '新建',
            add: {
                submit: '提交',
                cancel: '取消'
            }
        },
        footer: {
            one: {
                left: '由',
                right: '强力驱动'
            },
            two: {
                left: '已提供过',
                right: '次咨询'
            },
            three: "联系我们"
        }
    }, advice: {
        header: {
            title: '咨询建议',
            add: {
                title: '新建咨询',
                form: [
                    {
                        label: '咨询类型',
                        placeholder: 'e.g. 设计方案、会议纪要'
                    },
                    {
                        label: '描述',
                        placeholder: 'e.g. 修复了 XXX 问题'
                    }
                ],
            }
        }
    }, article: {
        header: {
            title: '文章 Q&A',
            errMsg: '标题不能为空 | 内容不能为空',
            add: {
                title: '新增文章',
                form: [
                    {
                        label: '文章标题',
                        placeholder: '',
                    },
                    {
                        label: '内容',
                        placeholder: '',
                    }
                ]
            }
        }, itemFrom: {
            placeholder: 'e.g. 请总结一下文章讲了什么？',
            submit: '新增提问',
            clear: '清空上下文',
            expendText: ['展开全文', '收起全文'],
        }
    }, settings: {
        title: '设置',
        first: {
            title: '应用设置',
            language: '语言',
            ownToken: '自备 Token',
        }, second: {
            title: '关于',
            project: '项目地址',
        }
    },
    empty: '暂无数据'
}

export default ZH
