export const certdResources = [
    {
        title: "证书自动化",
        name: "CertdRoot",
        path: "/certd",
        redirect: "/certd/pipeline",
        meta: {
            icon: "ion:key-outline",
            auth: true
        },
        children: [
            {
                title: "OpenKey",
                name: "OpenKey",
                path: "/certd/open/openkey",
                component: "/certd/open/openkey/index.vue",
                meta: {
                    icon: "hugeicons:api",
                    auth: true,
                    cache: true
                }
            },
            {
                title: "账号信息",
                name: "UserProfile",
                path: "/certd/mine/user-profile",
                component: "/certd/mine/user-profile.vue",
                meta: {
                    icon: "ion:person-outline",
                    auth: true,
                    isMenu: false
                }
            }
        ]
    }
];
