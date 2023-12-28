
<script>
import Post from "../components/Post.vue";
import { getCookie } from "../utils/cookieUtils";

export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			feed: null,
		};
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
                const identifier = getCookie('identifier');
                
                // this.$axios.defaults.headers.common['Authorization'] = identifier;
				let response = await this.$axios.get("/user/stream", { headers: { 'Authorization': identifier } })
				this.feed = response.data;
				console.log(response.data);
			}
			catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh();
	},
	components: { Post }
}
</script>

<template>
	<AuthWrapper :redirects="true">

		<div v-if="feed != null">

			<div class="feed" v-for="post in feed" :key="post.ID">
				<Post :post="post" />
			</div>

		</div>

	</AuthWrapper>
</template>
  