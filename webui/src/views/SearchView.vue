<script>
export default {
	data: function() {
		return {
			users: [],
		}
	},
	methods: {
		async searchUserByUsername() {

            let username = document.getElementById("username").value;
            try {
                let response = await this.$axios.get("/users", {params: {username: username}});
                this.users = response.data;
            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status ${status}: ${reason}`;
            }
			
		},
        visitProfile(Uid) {    
			this.$router.push(`/users/${Uid}/profile`);
		},
	},
}
</script>
<template>
    <div class="search-container">
        <form @submit.prevent="searchUserByUsername">
            <input id="username" class="search-box" type="text" placeholder="Username">
            <button class="search-button" type="submit">Search</button>
        </form>
    </div>

    <div v-for="user in users" :key="user.Uid">
        <div class="modal-body">
            <div class="container-fluid">
                <div class="row mb-2 mt-2">
                    <div class="col d-flex justify-content-center">
                        <div class="user-mini-card card bg-transparent border-start">
                            <div class="card-body">
                                <h5 @click="visitProfile(user.Uid)" class="user-mini-card-title d-flex justify-content-center ">@{{ user.username }}</h5>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

