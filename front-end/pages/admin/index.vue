<template>
<div class="margin__card">
    <v-row justify="center">
        <v-col xl="4" lg="4" md="6" sm="8" xs="12">
            <v-card>
                <v-card-title>
                    <h2>Login</h2>
                </v-card-title>
                <v-card-actions>
                    <form action="" class="form">
                        <v-text-field label="Enter Name" v-model="login.name" class="form__control" hide-details="auto">
                        </v-text-field>
                        <v-text-field @keydown.enter.prevent="loginBtn" type="password" label="Enter Password" v-model="login.password" class="form__control" hide-details="auto">
                        </v-text-field>
                        <v-btn color="error" @click="loginBtn" class="login__btn">Login</v-btn>
                    </form>
                </v-card-actions>
            </v-card>
        </v-col>
    </v-row>
</div>
</template>

<script>
export default {
    data() {
        return {
            login: {
                name: 'admin',
                password: '1',
            }
        }
    },
    methods: {
        loginBtn() {
            if (this.login.name.trim().length === 0) {
                this.$toast.warning("Veuillez saisir le nom !");
            } else if (this.login.password.trim().length === 0) {
                this.$toast.warning("Veuillez saisir le mot de passe !");
            } else {
                this.$auth.loginWith('local', {
                    data: {
                        action: 'login',
                        login: this.login
                    }
                }).then((data) => {
                    if (this.$auth.loggedIn === true) {
                        this.$auth.setUser(data.data.data.name)

                        this.$toast.success("Welcome !");
                    } else {
                        this.$toast.error("Erreur !");
                    }
                })
            }
        }
    }
};
</script>

<style>
</style>
