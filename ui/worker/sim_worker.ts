import './wasm_exec';
import { WorkerInterface } from './worker_interface';

type SimRequestAsync = (data: Uint8Array, progress: (result: Uint8Array) => void, id: string) => Uint8Array;
type SimRequestSync = (data: Uint8Array) => Uint8Array;

declare global {
	function wasmready(): void;
	const bulkSimAsync: SimRequestAsync;
	const computeStats: SimRequestSync;
	const computeStatsJson: SimRequestSync;
	const raidSim: SimRequestSync;
	const raidSimJson: SimRequestSync;
	const raidSimAsync: SimRequestAsync;
	const statWeights: SimRequestSync;
	const statWeightsAsync: SimRequestAsync;
	const statWeightRequests: SimRequestSync;
	const statWeightCompute: SimRequestSync;
	const raidSimRequestSplit: SimRequestSync;
	const raidSimResultCombination: SimRequestSync;
	const abortById: SimRequestSync;
}

globalThis.wasmready = function () {
	new WorkerInterface({
		bulkSimAsync: bulkSimAsync,
		computeStats: computeStats,
		computeStatsJson: computeStatsJson,
		raidSim: raidSim,
		raidSimJson: raidSimJson,
		raidSimAsync: raidSimAsync,
		statWeights: statWeights,
		statWeightsAsync: statWeightsAsync,
		statWeightRequests: statWeightRequests,
		statWeightCompute: statWeightCompute,
		raidSimRequestSplit: raidSimRequestSplit,
		raidSimResultCombination: raidSimResultCombination,
		abortById: abortById,
	}).ready(true);
};

const go = new Go();
WebAssembly.instantiateStreaming(fetch('/classic/lib.wasm'), go.importObject).then(async result => {
	await go.run(result.instance);
});

export {};