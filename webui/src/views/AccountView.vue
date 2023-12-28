
<script>
import Post from "../components/Post.vue";
import { useRoute } from 'vue-router';

export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			user_profile: null,
		};
	},
	setup() {
		const route = useRoute();
		const userId = route.params.userId;

		return { userId };
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
            const identifier = getCookie('identifier');
            
			try {
				let response = await this.$axios.get("/user/profile/"+identifier);
				this.user_profile = response.data;
				console.log(response.data);
			}
			catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		console.log("userId" + this.userId)
		this.refresh();
	},
	components: { Post }
}
</script>

<template>
	<AuthWrapper :redirects="true">

		<div v-if="user_profile != null">
			<!-- User Name -->
			<div class="header">
				<div class="top">
					<h1>{{ user_profile.Name }}</h1>
					<button class="follow-btn">Follow</button>
				</div>
				<div class="bottom">
					<h4>{{ user_profile.Following.length + " following" }}</h4>
					<h4>{{ user_profile.Followers.length + ((user_profile.Followers.length > 1) ? " followers" : " follower"
					) }}</h4>

				</div>
			</div>



			<div class="feed" v-for="post in user_profile.Photos" :key="post.ID">
				<Post :post="post" />
			</div>

		</div>

	</AuthWrapper>
</template>
  

  
 
  

<!-- <template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style> -->
