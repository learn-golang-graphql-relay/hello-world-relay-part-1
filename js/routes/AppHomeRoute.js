import Relay from 'react-relay';

// Your Relay routes
// Define a root GraphQL query into which your
// containers' query fragments will be composed.
export default class extends Relay.Route {
  static queries = {
    latestPost: () => Relay.QL`
      query {
        latestPost
      }
    `,
  };
  static routeName = 'AppHomeRoute';
}
