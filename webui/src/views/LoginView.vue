<template>
   <div class="login-container">
    <form @submit.prevent="doLogin">
      <div class="form-group">
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" required>
      </div>
      <button type="submit" :disabled="loading">Login</button>
    </form>
    <div v-if="loading">Logging in...</div>
    <div v-if="errormsg" class="error-message">{{ errormsg }}</div>
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