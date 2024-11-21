/* tslint:disable */
/* eslint-disable */
/**
 * tzetypes-badminton
 * tzetypes-badminton
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SignupClubOwnerRequestSchema
 */
export interface SignupClubOwnerRequestSchema {
    /**
     * 
     * @type {string}
     * @memberof SignupClubOwnerRequestSchema
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof SignupClubOwnerRequestSchema
     */
    password: string;
    /**
     * 
     * @type {string}
     * @memberof SignupClubOwnerRequestSchema
     */
    passwordRepeat: string;
}

/**
 * Check if a given object implements the SignupClubOwnerRequestSchema interface.
 */
export function instanceOfSignupClubOwnerRequestSchema(value: object): value is SignupClubOwnerRequestSchema {
    if (!('email' in value) || value['email'] === undefined) return false;
    if (!('password' in value) || value['password'] === undefined) return false;
    if (!('passwordRepeat' in value) || value['passwordRepeat'] === undefined) return false;
    return true;
}

export function SignupClubOwnerRequestSchemaFromJSON(json: any): SignupClubOwnerRequestSchema {
    return SignupClubOwnerRequestSchemaFromJSONTyped(json, false);
}

export function SignupClubOwnerRequestSchemaFromJSONTyped(json: any, ignoreDiscriminator: boolean): SignupClubOwnerRequestSchema {
    if (json == null) {
        return json;
    }
    return {
        
        'email': json['email'],
        'password': json['password'],
        'passwordRepeat': json['password_repeat'],
    };
}

export function SignupClubOwnerRequestSchemaToJSON(value?: SignupClubOwnerRequestSchema | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'email': value['email'],
        'password': value['password'],
        'password_repeat': value['passwordRepeat'],
    };
}

