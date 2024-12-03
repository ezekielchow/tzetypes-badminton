import {
  BaseAPI,
  GameApi,
  Configuration as PublicConf,
  ResponseError as PublicResponseError,
  UsersApi as PublicUsersApi,
  ResponseError,
  type ApiResponse,
  type GetGame200Response,
  type GetGameRequest,
  type InitOverrideFunction,
  type LoginRequest,
  type LoginResponseSchema,
  type SignupPlayerRequest
} from '@/repositories/clients/public';

import { resetStores } from './store';

export class MyPublicApi extends BaseAPI {

  private signInPageUrl: string = '/login';
  private backendUrl: string;

  constructor(backendUrl: string) {
    super(); // Call to the base class constructor if needed.
    this.backendUrl = backendUrl;
  }

  getPublicConf() {
    return new PublicConf({
      basePath: `${this.backendUrl}`,
      credentials: "include"
    });
  }


  async refreshToken(): Promise<void> {

    const api = new PublicUsersApi(this.getPublicConf())

    try {
      const response = await api.refreshToken()
      localStorage.setItem('session_token', response.sessionToken);

    } catch (error) {
      throw new Error('Failed to refresh token');
    }
  }

  private async authenticatedRequest(requestFunction: () => Promise<ApiResponse<any>>): Promise<ApiResponse<any>> {
    try {
      return await requestFunction();
    } catch (error) {
      if (error instanceof ResponseError && error.response.status === 401) {
        try {
          // Refresh the token if the request is unauthorized
          await this.refreshToken();

          // Retry the request after refreshing the token
          return await requestFunction();
        } catch (refreshError) {
          console.error('Token refresh failed: ', refreshError);

          // Logout and redirect to the sign-in page on failure
          this.deleteSession();
          throw new Error('Session expired. Please log in again.'); // Optionally rethrow to indicate the session issue
        }
      } else {
        // Rethrow the error if it's not a 401 or a refresh failure
        throw error;
      }
    }
  }


  private async requestLoginApi(requestParameters: LoginRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<LoginResponseSchema> {
    const api = new PublicUsersApi(this.getPublicConf())

    try {
      return await api.login(requestParameters, initOverrides)
    } catch (error) {

      if (error instanceof PublicResponseError) {
        throw new PublicResponseError(error.response, error.message)
      }
      throw new Error('Failed to login');
    }
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async login(requestParameters: LoginRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<LoginResponseSchema> {
    return await this.requestLoginApi(requestParameters, initOverrides);
  }

  /**
   * Logs out the user and redirects to the sign-in page
   */
  deleteSession(): void {
    // Clear session storage or any other storage that holds your authentication tokens
    sessionStorage.clear()
    localStorage.clear()

    resetStores()

    // Redirect the user to the sign-in page
    window.location.href = this.signInPageUrl;
  }

  private async getGameRequest(requestParameters: GetGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetGame200Response>> {
    const api = new GameApi(this.getPublicConf())

    try {
      return await api.getGameRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get game');
    }
  }

  async getGame(requestParameters: GetGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<GetGame200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.getGameRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async requestSignupPlayer(requestParameters: SignupPlayerRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<void> {
    const api = new PublicUsersApi(this.getPublicConf())

    try {
      return await api.signupPlayer(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to signup player');
    }
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async signupPlayer(requestParameters: SignupPlayerRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<void> {
    return await this.requestSignupPlayer(requestParameters, initOverrides);
  }
}