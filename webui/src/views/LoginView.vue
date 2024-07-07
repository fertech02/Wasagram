<script>
export default {
    components: {},

    data: function() {
        return { 
            errormsg: null,
        }
    },

    methods: {
        async doLogin() {
            if (this.username == "") {
                this.errormsg = "Please insert a username.";
                return;
            } else {
                try {
                    let username = document.getElementById('username').value;
				    if (!username.match("^[a-zA-Z][a-zA-Z0-9_]{2,15}$")) {
                        alert("Invalid username: 3 - 16 characters; first character must be a letter; only letters, numbers and underscores allowed");
                        return;
				    }
                    let response = await this.$axios.post("/session", {username: username},
                        {headers: { 'Content-Type': 'application/json' }});
                    let user = response.data
                    sessionStorage.setItem("token", user.Uid);
                    sessionStorage.setItem("username", user.Username);
                    this.$router.push({path: '/session'})
                } catch (error) {
                    if (error.response && error.response.data) {
                        this.errormsg = error.response.data.message;
                    } else {
                        // Handle cases where the error does not come from the server
                        this.errormsg = "An error occurred. Please try again.";
                    }
                }
            } 
        },
    },
    mounted() {
        
    }
}

</script>

<template>
    <div>
        <h2 class="h2">Login</h2>
        <div class="input-group">
            <input id="username" type="text" class="form-control" placeholder="Username" required>
            <button class="btn btn-success" @click="doLogin">Login</button>
        </div>
        <div v-if="errormsg" class="alert alert-danger" role="alert">
            {{ errormsg }}
        </div>
    </div>
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