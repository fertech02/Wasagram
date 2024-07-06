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

<script>
export default {
    data: function() {
        return {
            username: '',
            identifier: '',
            loading: false,
            errormsg: null,

        }
    },
    methods: {
        async doLogin() {
            this.loading = true;
            this.errormsg = null;

            try {
                let response = await this.$axios.post("/session", {username: this.username},{
                    headers: {
                        'Accept' : 'application/json',
                        'Content-Type' : 'application/json'
                    },
                });
                this.identifier = response.data.identifier;
                localStorage.setItem("token", this.identifier);
                localStorage.setItem("username", this.username);
                this.$router.push({path: '/session'});
            } catch (error) {
                this.errormsg = error.response.data.message;
            } 
            this.loading = false;
        },
    }
}
</script>

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