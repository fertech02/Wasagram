<script>
export default {
    data(){
    return{
            token: sessionStorage.getItem('token'),
            username: "",
        }		
    },
    methods: {
        async setMyUsername() {
			try {
				let username = document.getElementById('username').value;
				let response = await this.$axios.put("/users/" + this.token + "/username", {username: username});
				let user = response.data;
				sessionStorage.setItem('username', user.Username);
			} catch (error) {
				console.log(error);
        }
    }
	},
}
</script>

<template>
	<div>
		<h1>Set My Username</h1>
		<input type="text" id="username" placeholder="Enter your new username">
		<button @click="setMyUsername">Set Username</button>
	</div>
</template>