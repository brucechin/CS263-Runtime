echo "[Rust][Benchmark] http server"

echo
echo "run command 'wrk --latency -t4 -c200 -d8s http://127.0.0.1:8082' from another terminal"
echo "Please manually stop this program when wrk finishes reporting"
cd hyper && cargo build --release

cd ..
echo