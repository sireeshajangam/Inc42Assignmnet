// nextjs-frontend/pages/index.tsx

import React from 'react';
import styled from 'styled-components';

const Container = styled.div`
	text-align: center;
	padding: 50px;
	background-color: #f0f0f0;
`;

const Button = styled.button`
	padding: 10px 20px;
	font-size: 16px;
	background-color: #007bff;
	color: #fff;
	border: none;
	border-radius: 5px;
	cursor: pointer;

	&:hover {
		background-color: #0056b3;
	}
`;

const Home: React.FC = () => {
	return (
		<Container>
			<h1>Hello from Next.js Frontend!</h1>
			<Button>Click Me</Button>
		</Container>
	);
};

export default Home;
