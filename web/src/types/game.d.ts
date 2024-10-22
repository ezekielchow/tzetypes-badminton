

export type Player = {
  id: string
  name: string
}

export type Step = {
  scoreAt: Date
  teamLeftScore: number
  teamRightScore: number
  currentServer: CurrentServer
  isEnded: boolean
}

export type GameState = {
  type?: GameTypes
  leftOddPlayer?: Player
  leftEvenPlayer: Player
  rightOddPlayer?: Player
  rightEvenPlayer: Player
  firstServer: CurrentServer
  progress: [Step?]
}

