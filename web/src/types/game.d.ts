import type { GameStep } from "@/repositories/clients/private";

interface LocalGameStep extends GameStep {
  isSynced: boolean,
}