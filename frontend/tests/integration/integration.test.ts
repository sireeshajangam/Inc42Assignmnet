// Sample integration test using Jest
import { fetchData } from '../index';

test('fetchData function returns posts', async () => {
  const data = await fetchData();
  expect(data.length).toBeGreaterThan(0);
});
