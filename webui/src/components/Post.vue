<script>

function userHasLiked(post, identifier) {
    return post.Likes.some(like => like.UserID == identifier);
}

export default {
    props: ["post", "identifier"],
    data: function () {
        return {
            liked: userHasLiked(this.post, this.identifier),
            newComment: ''
        };
    },
    methods: {
        deleteImage() {
            window.confirm("Are you sure you want to delete this image?") &&
                this.$axios.delete("/photos?photoId=" + this.post.ID, { headers: { 'Authorization': this.identifier } })
                    .then((response) => {
                        console.log(response);
                        this.$emit('delete', this.post.ID);
                    })
                    .catch((error) => {
                        console.log(error);
                    })
        },
        toggleLike() {
            console.log(this.identifier)
            if (this.liked) {
                this.$axios.delete("/photos/" + this.post.ID+ "/likes" , { headers: { 'Authorization': this.identifier } })
                    .then((response) => {
                        console.log(response);
                        this.post.Likes = response.data;
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
            else {
                this.$axios.post("/photos/" + this.post.ID+ "/likes", {}, { headers: { 'Authorization': this.identifier } })
                    .then((response) => {
                        console.log(response);
                        this.post.Likes = response.data;
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
        },
        submitComment() {
            const payload = {
                comment: this.newComment
            };
            this.$axios.post('/comments?photoId=' + this.post.ID, payload, {
                headers: { 'Authorization': this.identifier }
            })
                .then((response) => {
                    console.log(response);
                    this.post.Comments.push(response.data);
                    this.newComment = ''; // Reset comment input
                })
                .catch((error) => {
                    console.log(error);
                });
        },
        deleteComment(id) {
            console.log(this.post)
            this.$axios.delete('/comments?commentId=' + id, {
                headers: { 'Authorization': this.identifier }
            })
                .then((response) => {
                    console.log(response);
                    this.post.Comments = this.post.Comments.filter(comment => comment.ID != id);
                })
                .catch((error) => {
                    console.log(error);
                });
        }
    }
}
</script>

<template>
    <div v-if="post != null" class="post-container">
        <div class="username-container">
            <RouterLink :to="{ path: 'user/' + post.UserID }" replace>
                {{ post.UserName }}
                <!-- <p>{{ this.test}}</p> -->
            </RouterLink>
            <button class="" @click="deleteImage" v-if="post.UserID == this.identifier">
                <svg class="feather" :class="this.liked ? 'liked' : ''">
                    <use href="/feather-sprite-v4.29.0.svg#trash" />
                </svg>
            </button>
        </div>
        <div class="image-container">
            <img :src="'http://localhost:3000/images/' + post.FileName" :alt="post.FileName">
        </div>
        <div class="likes-container">
            <span>
                <button class="like-btn" @click="toggleLike">
                    <svg class="feather" :class="this.liked ? 'liked' : ''">
                        <use href="/feather-sprite-v4.29.0.svg#heart" />
                    </svg>
                </button>
                {{ post.Likes.length + ((post.Likes.length > 1) ? " likes" : " like") }}
            </span>
            <span class="date">
                <!--  display the time with month day year and only the hour-->
                {{ new Date(post.UploadTime).toLocaleString('default', {
                    month: 'long',
                    year: 'numeric',
                    hour: '2-digit',
                    hour12: true // change to false if you want 24-hour format
                }) }}
                <!-- {{ new Date(post.UploadTime).toLocaleDateString()  }} -->
            </span>
        </div>
        <div class="comments-container">
            <ul>
                <li class="comment" v-for="comment in post.Comments" :key="comment.ID">
                    <div class="comment-top">

                        <div>
                            <strong>{{ comment.UserName }}</strong> - <small>{{ new
                                Date(comment.CreateTime).toLocaleString()
                            }}</small>
                        </div>
                        <button class="" @click="() => deleteComment(comment.ID)" v-if="comment.UserID == this.identifier">
                            <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#trash" />
                            </svg>
                        </button>
                    </div>
                    {{ comment.Text }}

                </li>
            </ul>
            <form class="comment-input-container" @submit.prevent="submitComment">
                <input type="text" v-model="newComment" placeholder="Comment" />
                <button type="submit">Post</button>
            </form>
        </div>
    </div>
</template>

<style></style>