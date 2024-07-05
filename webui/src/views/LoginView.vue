<template>
    <div class="login-view">
        <h1>Login</h1>
        <LoadingSpinner v-if="loading" ></LoadingSpinner>
        <div class="login-form">
            <form @submit.prevent="login">
                <input type="text" id="username" v-model="username"  required minlength="3" maxlength="16" style="padding: 8px;"/>
                <button type="submit" class="btn btn-sm btn-outline-primary" style="padding: 8px; float: right; font-size: large;" >Login <svg class="feather"> <use href="/feather-sprite-v4.29.0.svg#key" /></svg></button>
            </form>
            <div v-if="identifier != null">
                <p>Logged in as: {{identifier}}</p>
            </div>
        </div>
    </div>
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
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f8f9fa;
}

.login-form {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    background-color: #fff;
    border-radius: 4px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.login-label {
    padding: 3px;
    display: block;
    margin-bottom: 8px;
}
</style>