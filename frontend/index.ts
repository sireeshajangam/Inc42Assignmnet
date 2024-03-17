async function fetchPosts() {
    const response = await fetch('http://localhost:8080/posts');
    const posts = await response.json();
    return posts;
  }
  
  async function displayPosts() {
    const posts = await fetchPosts();
    const postList = document.getElementById('post-list');
    posts.forEach((post: { title: string; body: string }) => {
      const listItem = document.createElement('li');
      listItem.innerHTML = `<h2>${post.title}</h2><p>${post.body}</p>`;
      postList?.appendChild(listItem);
    });
  }
  
  displayPosts();
  