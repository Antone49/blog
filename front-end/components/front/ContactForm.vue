<template>
<div class="my-10">
    <v-container>
        <h1 class="m-3">Contact Us</h1>
        <v-card>
            <v-card-title>
                <v-row>
                    <v-col xl="12" lg="12" md="12" sm="12">
                        <form action="" class="form">
                            <v-text-field label="Email" v-model="email" class="form__control my-3" hide-details="auto" outlined>
                            </v-text-field>
                            <v-text-field label="Subject" v-model="subject" class="form__control my-3" hide-details="auto" outlined>
                            </v-text-field>
                            <v-textarea class="form__control" label="Message" v-model="message" outlined></v-textarea>
                            <v-btn color="error" @click="contactUsBtn" class="login__btn">Contact Us</v-btn>
                        </form>
                    </v-col>
                </v-row>
            </v-card-title>
        </v-card>
    </v-container>
</div>
</template>

<script>
import {
    contactUs
} from '../../functions/mail.js'

export default {
    name: "ContactVue",
    data() {
        return {
            alert: true,
            email: '',
            subject: '',
            message: ''
        }
    },
    methods: {
        contactUs
    },
    methods: {
        contactUsBtn() {
            if (this.email.trim().length === 0) {
                this.$toast.warning("Veuillez saisir l'adresse mail !");
            } else if (this.subject.trim().length === 0) {
                this.$toast.warning("Veuillez saisir le sujet !");
            } else if (this.message.trim().length === 0) {
                this.$toast.warning("Veuillez saisir le message !");
            } else {
                const regex = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                const found = this.email.match(regex);

                if (found === null) {
                    this.$toast.warning("Adresse mail incorrecte !");
                } else {
                    contactUs(this.email, this.subject, this.message)
                        .then((data) => {
                            if (data === true) {
                                this.email = ""
                                this.subject = ""
                                this.message = ""

                                this.$toast.success("Message envoyé !");
                            } else {

                                this.$toast.error("Erreur pour envoyer le message !");
                            }
                        })
                }
            }
        }
    }
};
</script>

<style >
</style>
