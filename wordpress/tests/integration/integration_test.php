<?php
// Sample integration test for WordPress functions using PHPUnit
class IntegrationTest extends WP_UnitTestCase {
    public function test_get_custom_posts_integration() {
        // Add test posts
        $posts = get_custom_posts();

        $this->assertEquals(2, count($posts));
        $this->assertEquals('First Post', $posts[0]['title']);
        $this->assertEquals('Second Post', $posts[1]['title']);
    }
}
