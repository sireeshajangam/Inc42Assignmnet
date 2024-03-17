<?php
// Assuming this code is within a WordPress theme or plugin

function get_custom_posts() {
  $args = array(
    'post_type' => 'post',
    'posts_per_page' => -1,
  );
  $query = new WP_Query($args);
  $posts = array();
  if ($query->have_posts()) {
    while ($query->have_posts()) {
      $query->the_post();
      $post_data = array(
        'id' => get_the_ID(),
        'title' => get_the_title(),
        'body' => get_the_content(),
      );
      array_push($posts, $post_data);
    }
    wp_reset_postdata();
  }
  return $posts;
}
