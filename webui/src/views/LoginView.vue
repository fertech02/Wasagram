<script>

export default {
    components: {},

    data: function() {
        return {
            username: "",
            errormsg: null,
            Profile: {
                username: '',
                identifier: ''
            }

        }
    },

    methods: {
        async doLogin() {
            this.errormsg = null;
            if (this.username == "") {
                this.errormsg = "Please insert a username.";
                return;
            }
            try {
                let response = await this.$axios.post("/session", {username: this.username});
                this.Profile = response.data;
                localStorage.setItem("token", this.Profile.identifier);
                localStorage.setItem("username", this.Profile.username);
                this.$router.push({path: '/session'});
            } catch (error) {
                if (error.response && error.response.data) {
                     this.errormsg = error.response.data.message;
                } else {
                    // Handle cases where the error does not come from the server
                    this.errormsg = "An error occurred. Please try again.";
                }
            } 
        },
    }
}

</script>

<template>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Welcome to Wasaphoto</h1>
    </div>
    <div class="input-group mb-3">
        <input type="text" id="username" v-model="username" class="form-control"
            placeholder="Insert a username to log in Wasaphoto." aria-label="Recipient's username"
            aria-describedby="basic-addon2">
        <div class="input-group-append">
            <button class="btn btn-success" type="button" @click="doLogin">Login</button>
        </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>


<style scoped>

    .input-group {
        width: 50%;
        margin: auto;
    }

    .btn-success {
        width: 100%;
    }

    .h2 {
        margin: auto;
    }

</style>