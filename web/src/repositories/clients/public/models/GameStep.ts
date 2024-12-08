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
 * @interface GameStep
 */
export interface GameStep {
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    gameId: string;
    /**
     * 
     * @type {number}
     * @memberof GameStep
     */
    teamLeftScore: number;
    /**
     * 
     * @type {number}
     * @memberof GameStep
     */
    teamRightScore: number;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    scoreAt: string;
    /**
     * 
     * @type {number}
     * @memberof GameStep
     */
    stepNum: number;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    currentServer: string;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    leftOddPlayerName: string;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    leftEvenPlayerName: string;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    rightOddPlayerName: string;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    rightEvenPlayerName: string;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    syncId?: string;
    /**
     * 
     * @type {number}
     * @memberof GameStep
     */
    isPaused: number;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof GameStep
     */
    updatedAt: string;
}

/**
 * Check if a given object implements the GameStep interface.
 */
export function instanceOfGameStep(value: object): value is GameStep {
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('gameId' in value) || value['gameId'] === undefined) return false;
    if (!('teamLeftScore' in value) || value['teamLeftScore'] === undefined) return false;
    if (!('teamRightScore' in value) || value['teamRightScore'] === undefined) return false;
    if (!('scoreAt' in value) || value['scoreAt'] === undefined) return false;
    if (!('stepNum' in value) || value['stepNum'] === undefined) return false;
    if (!('currentServer' in value) || value['currentServer'] === undefined) return false;
    if (!('leftOddPlayerName' in value) || value['leftOddPlayerName'] === undefined) return false;
    if (!('leftEvenPlayerName' in value) || value['leftEvenPlayerName'] === undefined) return false;
    if (!('rightOddPlayerName' in value) || value['rightOddPlayerName'] === undefined) return false;
    if (!('rightEvenPlayerName' in value) || value['rightEvenPlayerName'] === undefined) return false;
    if (!('isPaused' in value) || value['isPaused'] === undefined) return false;
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    return true;
}

export function GameStepFromJSON(json: any): GameStep {
    return GameStepFromJSONTyped(json, false);
}

export function GameStepFromJSONTyped(json: any, ignoreDiscriminator: boolean): GameStep {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'],
        'gameId': json['game_id'],
        'teamLeftScore': json['team_left_score'],
        'teamRightScore': json['team_right_score'],
        'scoreAt': json['score_at'],
        'stepNum': json['step_num'],
        'currentServer': json['current_server'],
        'leftOddPlayerName': json['left_odd_player_name'],
        'leftEvenPlayerName': json['left_even_player_name'],
        'rightOddPlayerName': json['right_odd_player_name'],
        'rightEvenPlayerName': json['right_even_player_name'],
        'syncId': json['sync_id'] == null ? undefined : json['sync_id'],
        'isPaused': json['is_paused'],
        'createdAt': json['created_at'],
        'updatedAt': json['updated_at'],
    };
}

export function GameStepToJSON(value?: GameStep | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'game_id': value['gameId'],
        'team_left_score': value['teamLeftScore'],
        'team_right_score': value['teamRightScore'],
        'score_at': value['scoreAt'],
        'step_num': value['stepNum'],
        'current_server': value['currentServer'],
        'left_odd_player_name': value['leftOddPlayerName'],
        'left_even_player_name': value['leftEvenPlayerName'],
        'right_odd_player_name': value['rightOddPlayerName'],
        'right_even_player_name': value['rightEvenPlayerName'],
        'sync_id': value['syncId'],
        'is_paused': value['isPaused'],
        'created_at': value['createdAt'],
        'updated_at': value['updatedAt'],
    };
}

