<template>
<div>
    <v-container>
        <v-row>
            <v-col xl="8" lg="8" md="8" sm="12">
                <v-btn color="error" class="login__btn">J'aime</v-btn>
                <h3>Comments( 44 )</h3>
                <hr />
            </v-col>
        </v-row>
        <v-row>
            <v-col xl="7" lg="7" md="7" sm="12">
                <div v-for="i in 6" :key="i">
                    <v-sheet>
                        <h3>Hilal ahmad</h3>
                        <p>comment</p>
                        <hr />
                    </v-sheet>
                </div>
            </v-col>
            <v-col xl="5" lg="5" md="5" sm="12" xs="12">
                <form action="" class="form">
                    <v-text-field label="Enter Your Name" class="form__control" hide-details="auto">
                    </v-text-field>
                    <v-text-field label="Enter Your Email" class="form__control" hide-details="auto">
                    </v-text-field>
                    <v-text-field label="Enter Your Comment" class="form__control" hide-details="auto">
                    </v-text-field>
                    <v-btn color="error" class="login__btn"> Comment</v-btn>
                </form>
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
export default {
    name: "PostMessageComp",
    props: {

    },
};
</script>

<style >
</style>
