import { useState } from "react";
import { RunSimpleSimulation } from "../wailsjs/go/main/App";
import { statistics } from "../wailsjs/go/models";
import {
  Area,
  AreaChart,
  CartesianGrid,
  Line,
  LineChart,
  ResponsiveContainer,
  XAxis,
  YAxis,
} from "recharts";

function App() {
  const [res, setRes] = useState<statistics.LearnedSummary[]>();
  const [precision, setPrecision] = useState<number>(3); // Initial value for precision
  const [yearCount, setYearCount] = useState<number>(1); // Initial value for years
  const [principal, setPrincipal] = useState<number>(1000); // Initial value for years
  const [additional, setAdditional] = useState<number>(1000); // Initial value for years
  const [investment, setInvestment] = useState<string>("market"); // Initial value for years
  const [isLoading, setIsLoading] = useState<boolean>(false); // Loading state

  const precisionOptions = [0.01, 0.1, 1, 10, 100, 1000];

  const handlePrecisionChange = (event: any) => {
    setPrecision(event.target.value);
  };

  function doRunSimpleSimulation() {
    setIsLoading(true); // Set loading to true when starting the simulation
    RunSimpleSimulation(
      precisionOptions[precision],
      yearCount,
      principal,
      investment,
      additional
    ).then((result) => {
      setRes(result);
      setIsLoading(false); // Set loading to false when simulation is done
    });
  }

  return (
    <div className="font-mono w-full h-[100vh] bg-slate-200 text-slate-900 flex flex-row">
      <div className="flex flex-col w-1/3 border-r-2 border-slate-400 h-100% justify-between">
        <div className="flex flex-col justify-start">
          <h1 className="text-lg divide-x-2 pt-2 pb-2 border-b-2 border-slate-500">
            Configure Simulation
          </h1>
          <div className="flex flex-col border-b-2 pt-2 pb-2 border-slate-500">
            <h2 className="pb-2">
              precise down to {precisionOptions[precision]} dollars
            </h2>
            <input
              className="h-2 pb-0 mb-0 mx-6 bg-slate-500 rounded-lg appearance-none cursor-pointer text-green-500"
              type="range"
              min="0"
              max="5"
              step="1"
              value={precision}
              onChange={handlePrecisionChange}
            />
            <div className="flex flex-row justify-between m-4">
              <h3>{0.01}</h3>
              <h3>{1000}</h3>
            </div>
          </div>
          <div className="flex flex-col align-top text-start pt-2 pb-2 pl-4 border-b-2 border-slate-500">
            <h3>Principal</h3>
            <input
              className="w-5/6 h-6 pb-0 mb-0 bg-slate-300 rounded-sm mt-2"
              type="number"
              min="0"
              max="1_000_000"
              step="1"
              value={principal}
              onChange={(e) => setPrincipal(parseFloat(e.target.value))}
            />
          </div>
          <div className="flex flex-col align-top text-start pt-2 pb-2 pl-4 border-b-2 border-slate-500">
            <h3>Additional Contribution</h3>
            <input
              className="w-5/6 h-6 pb-0 mb-0 bg-slate-300 rounded-sm mt-2"
              type="number"
              min="0"
              max="1_000_000"
              step="1"
              value={additional}
              onChange={(e) => setAdditional(parseFloat(e.target.value))}
            />
          </div>
          <div className="flex flex-col align-top text-start pt-2 pb-2 pl-4 border-b-2 border-slate-500">
            <h3>How Many Years</h3>
            <input
              className="w-5/6 h-6 pb-0 mb-0 bg-slate-300 rounded-sm mt-2"
              type="number"
              min="0"
              max="150"
              step="1"
              value={yearCount}
              onChange={(e) => setYearCount(parseFloat(e.target.value))}
            />
          </div>
          <div className="flex flex-col align-top text-start pt-2 pb-2 pl-4 border-b-2 border-slate-500">
            <h3>Economic Factor</h3>
            <select
              className="w-5/6 h-6 pb-0 mb-0 bg-slate-300 rounded-sm mt-2"
              value={investment}
              onChange={(e) => setInvestment(e.target.value)}
            >
              <option value="">Select an option</option>
              <option value="market">Market</option>
              <option value="bonds">Bonds</option>
            </select>
          </div>
        </div>
        <div className="flex flex-col justify-end">
          <button
            className="bg-green-500 rounded-lg h-10 m-4 flex flex-col justify-center pl-8"
            onClick={() => doRunSimpleSimulation()}
            disabled={isLoading} // Disable the button when loading
          >
            {isLoading ? (
              <div className="justify-center animate-spin border-2 border-slate-900 border-t-transparent w-5 h-5 rounded-full"></div>
            ) : (
              "Simulate"
            )}
          </button>
        </div>
      </div>
      <div className="w-2/3 h-full pt-8 flex flex-col">
        <ResponsiveContainer width="90%" height="50%">
          <AreaChart
            data={
              res?.map((item, index) => ({
                name: `${index + 1}`,
                q2: [item.Q2, item.Q2],
                range: [item.Q1, item.Q3],
              })) || []
            }
          >
            <CartesianGrid stroke="#ccc" />
            <XAxis dataKey="name" />
            <YAxis />
            {/* Shaded area between Q1 and Q3 */}
            <Area
              animationDuration={100}
              dataKey="range"
              fill="#4ade80"
              strokeWidth={0}
              fillOpacity={0.3}
            />

            {/* Line for Q2 (median) */}
            <Area
              animationDuration={100}
              dataKey="q2"
              stroke="#16a34a"
              strokeWidth={4}
              dot={{
                stroke: "#16a34a",
                strokeWidth: 3,
                r: 3,
                strokeDasharray: "",
              }}
            />
          </AreaChart>
        </ResponsiveContainer>
        <ResponsiveContainer width="90%" height="50%">
          <LineChart
            data={
              res?.map((item, index) => ({
                name: `${index + 1}`,
                mean: item.PPF,
              })) || []
            }
          >
            <Line type="monotone" dataKey="mean" stroke="#8884d8" />
            <CartesianGrid stroke="#ccc" />
            <XAxis dataKey="name" />
            <YAxis max={1} min={0} />
          </LineChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}

// TODO:
// take multiple investments at the same time
// include more than exist right now like real estate or gold
// use websockets to get more clear commmunication and stop early / see live results
export default App;
