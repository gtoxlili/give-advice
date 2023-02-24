const EN = {
    sideBar: {
        tabsTitle: ['Suggestions', 'Articles', 'Settings']
    }, inquiry: {
        header: {
            button: 'New',
            add: {
                submit: 'Submit',
                cancel: 'Cancel'
            }
        },
        footer: {
            one: {
                left: 'Powered by',
                right: 'Strong'
            },
            two: {
                left: 'Provided',
                right: 'Consultations'
            },
            three: "Contact Us"
        }
    }, advice: {
        header: {
            title: 'Advice',
            add: {
                title: 'New Advice',
                form: [
                    {
                        label: 'Type of advice',
                        placeholder: 'e.g. Design plan, Meeting minutes'
                    },
                    {
                        label: 'Description',
                        placeholder: 'e.g. Fixed the XXX problem'
                    }
                ],
            }
        }
    }, article: {
        header: {
            title: 'Articles Q&A',
            errMsg: 'Title cannot be empty | Content cannot be empty',
            add: {
                title: 'Add Article',
                form: [
                    {
                        label: 'Article Title',
                        placeholder: '',
                    },
                    {
                        label: 'Content',
                        placeholder: '',
                    }
                ]
            }
        }, itemFrom: {
            placeholder: 'e.g. Can you summarize what the article is about?',
            submit: 'Add Question',
            clear: 'Clear Context',
            expendText: ['Expand', 'Collapse'],
        }
    }, settings: {
        title: 'Settings',
        first: {
            title: 'App Settings',
            language: 'Language',
            ownToken: 'Own Token',
        }, second: {
            title: 'About',
            project: 'Project Address',
        }
    },
    empty: 'No Data'
}

export default EN
