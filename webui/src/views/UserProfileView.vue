
<script>
import Post from "../components/Post.vue";
import Modal from "../components/Modal.vue";
import { useRoute, useRouter } from 'vue-router';
import { getCookie } from "../utils/cookieUtils";
import { watch } from 'vue';
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			user_profile: null,
			followingModalIsVisible: false,
			followersModalIsVisible: false,
		};
	},
	setup() {
		const route = useRoute();
		const router = useRouter();
		const userId = route.params.userId;

		const identifier = getCookie('identifier');
		watch(() => route.params.userId, (newUserId) => {
			router.go(0);
		});
		return { userId, identifier };

	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/user/profile/" + this.userId);
				this.user_profile = response.data;
			}
			catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		handleFollowersToggle() {
			this.followersModalIsVisible = !this.followersModalIsVisible;
		},
		handleFollowingToggle() {
			this.followingModalIsVisible = !this.followingModalIsVisible;
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
				<Modal :show="followersModalIsVisible" @close="handleFollowersToggle" :users="user_profile.Followers">
					<template v-slot:header>
						<h3>Followers</h3>
					</template>
				</Modal>
				<Modal :show="followingModalIsVisible" @close="handleFollowingToggle" :users="user_profile.Following">
					<template v-slot:header>
						<h3>Following</h3>
					</template>
				</Modal>
				<div class="top">
					<h1>{{ user_profile.Name }}</h1>
					<button class="follow-btn">Follow</button>
				</div>
				<div class="bottom">
					<h4 @click="handleFollowingToggle">{{ user_profile.Following.length + " following" }}</h4>
					<h4 @click="handleFollowersToggle">{{ user_profile.Followers.length + ((user_profile.Followers.length >
						1) ? " following" : " follower"
					) }}</h4>
				</div>
			</div>



			<div class="feed" v-for="post in user_profile.Photos" :key="post.ID">
				<Post :post="post" :identifier="this.identifier" />
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
