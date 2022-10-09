<template>
<div>
    <v-container>
        <v-row>
            <v-col xl="8" lg="8" md="8" sm="12">
                <h3 v-if="commentaires != null">Commentaires ({{commentaires.length}})</h3>
                <h3 v-else>Pas de commentaires</h3>
                <hr />
            </v-col>
        </v-row>
        <v-row>
            <v-col xl="7" lg="7" md="7" sm="12">
                <div v-for="(item, index) in commentaires" :key="index">
                    <v-sheet>
                        <h3>{{ item.username }}</h3>
                        <p>{{ item.message }}</p>
                        <hr />
                    </v-sheet>
                </div>
            </v-col>
            <v-col xl="5" lg="5" md="5" sm="12" xs="12">
                <form action="" class="form">
                    <v-text-field label="Nom" class="form__control" hide-details="auto" v-model="username">
                    </v-text-field>
                    <v-text-field label="Commentaire" class="form__control" hide-details="auto" v-model="comment">
                    </v-text-field>
                    <v-text-field label="Email" class="form__control" hide-details="auto" v-model="email">
                    </v-text-field>
                    <v-btn color="error" class="login__btn" @click="commentBtn">Commenter</v-btn>
                </form>
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
import {
    addMessage,
    getAllMessagesFromPostId
} from '../../functions/message.js'

export default {
    name: "PostMessageComp",
    data() {
        return {
            username: 'rfr',
            comment: 'frf',
            email: 'fefe@frfr.ed',
            commentaires: null,
            postId: null
        }
    },
    methods: {
        commentBtn() {
            if (this.username.trim().length === 0) {
                this.$toast.warning("Veuillez saisir un nom !");
            } else if (this.comment.trim().length === 0) {
                this.$toast.warning("Veuillez saisir un commentaire !");
            } else if (this.email.trim().length === 0) {
                this.$toast.warning("Veuillez saisir votre email !");
            } else {
                const regex = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                const found = this.email.match(regex);

                if (found === null) {
                    this.$toast.warning("Adresse mail incorrecte !");
                } else {
                    addMessage(this.username, this.comment, this.email, this.postId)
                        .then((data) => {
                            if (data.error === false) {
                                this.username = ""
                                this.comment = ""

                                this.$toast.success("Commentaire envoyé !");
                            } else {
                                this.$toast.error("Erreur pour envoyer le commentaire !");
                            }
                        })
                }
            }
        }
    },
    mounted: function () {
        this.postId = parseInt(this.$route.query.id)

        getAllMessagesFromPostId(this.postId).then(result => {
            this.commentaires = result
        });
    },
};
</script>

<style >
</style>
