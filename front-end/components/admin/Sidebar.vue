<template>
<div>
    <v-navigation-drawer v-model="drawer" :mini-variant="miniVariant" :clipped="clipped" fixed app>
        <v-list>
            <v-list-item v-for="(item, i) in items" :key="i" :to="item.to" router exact>
                <v-list-item-action>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-list-item-action>
                <v-list-item-content>
                    <v-list-item-title v-text="item.title" />
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="clipped" flat app>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
        <v-toolbar-title v-text="title" />
        <v-spacer />
        <v-btn id="menu-activator" color="primary">Hi Admin</v-btn>

        <v-menu activator="#menu-activator">
            <v-list>
                <v-list-item>
                    <v-list-item-title>
                        <v-btn color="error" @click="logoutBtn" class="btn__width">Logout</v-btn>
                    </v-list-item-title>
                </v-list-item>
            </v-list>
        </v-menu>
    </v-app-bar>
</div>
</template>

<script>
export default {
    name: "SidebarComp",
    data() {
        return {
            clipped: false,
            drawer: true,
            fixed: false,
            items: [{
                    icon: "mdi-apps",
                    title: "Dashboard",
                    to: "/admin/dashboard",
                },
                {
                    icon: "mdi-post",
                    title: "Post",
                    to: "/admin/post",
                },
                {
                    icon: "mdi-widgets",
                    title: "Catégorie",
                    to: "/admin/category",
                },
                {
                    icon: "mdi-map-marker-outline",
                    title: "Location",
                    to: "/admin/location",
                },
                {
                    icon: "mdi-comment-outline",
                    title: "Comment",
                    to: "/admin/comment",
                },
            ],
            miniVariant: false,
            right: true,
            rightDrawer: false,
            title: "Blog",
        };
    },
    methods: {
        logoutBtn() {
            this.$auth.logout({
                data: {
                    action: 'logout',
                    logout: {
                      token: this.$auth.strategy.token.get(),
                      user: this.$auth.user,
                    }
                }
            }).then((data) => {
                console.log(this.$auth.loggedIn)
                console.log(this.$auth.strategy.token.get())

                if (this.$auth.loggedIn === false) {
                    this.$toast.success("Déconnecté !");
                } else {

                    this.$toast.error("Erreur !");
                }
            })
        }
    }
};
</script>

<style >
</style>
