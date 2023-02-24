const JA = {
    sideBar: {
        tabsTitle: ['提案', '記事', '設定']
    }, inquiry: {
        header: {
            button: '新規作成',
            add: {
                submit: '送信',
                cancel: 'キャンセル'
            }
        },
        footer: {
            one: {
                left: 'Powered by',
                right: ''
            },
            two: {
                left: 'すでに',
                right: '件の問い合わせ'
            },
            three: "お問い合わせ"
        }
    }, advice: {
        header: {
            title: 'アドバイスの提案',
            add: {
                title: '新しいアドバイスを作成する',
                form: [
                    {
                        label: 'アドバイスの種類',
                        placeholder: 'e.g. 設計案、会議の議事録'
                    },
                    {
                        label: '説明',
                        placeholder: 'e.g. XXX 問題が修正されました'
                    }
                ],
            }
        }
    }, article: {
        header: {
            title: '記事 Q&A',
            errMsg: 'タイトルが空白です | コンテンツが空白です',
            add: {
                title: '新しい記事を追加する',
                form: [
                    {
                        label: '記事タイトル',
                        placeholder: '',
                    },
                    {
                        label: '内容',
                        placeholder: '',
                    }
                ]
            }
        }, itemFrom: {
            placeholder: 'e.g. 記事の要約を書いてください',
            submit: '新しい質問を追加する',
            clear: '文脈をクリアする',
            expendText: ['全文を開く', '全文を閉じる'],
        }
    }, settings: {
        title: '設定',
        first: {
            title: 'アプリケーションの設定',
            language: '言語',
            ownToken: '自分でトークンを持つ',
        }, second: {
            title: 'このアプリについて',
            project: 'プロジェクトのアドレス',
        }
    },
    empty: 'データがありません'
}

export default JA
