import subprocess
import glob
import os
print("========== Generating GO files ==========")
GOPATH = os.environ.get('GOPATH')
if not GOPATH:
    raise Exception('GOPATH not set')

proto_files = glob.glob('proto/*.proto')

for f in proto_files:
    cmd = ['protoc', f'--go_out=plugins=grpc:{os.path.join(GOPATH,"src")}', f]
    print(' '.join(cmd))
    subprocess.run(cmd, shell=False)

print("========== Generating TS files ==========")
SCRIPT_PATH = os.path.dirname(os.path.abspath(__file__))
OUT_DIR = os.path.join(SCRIPT_PATH, 'ui', 'src', 'protogen')
os.makedirs(OUT_DIR, exist_ok=True)
PROTOC_GEN_TS_PATH = os.path.join(SCRIPT_PATH, 'ui', 'node_modules', '.bin', 'protoc-gen-ts.cmd' if os.name == 'nt' else 'protoc-gen-ts')
for f in proto_files:
    cmd = ['protoc',
           f'--plugin=protoc-gen-ts={PROTOC_GEN_TS_PATH}',
           f'--js_out=import_style=commonjs,binary:{OUT_DIR}',
           f'--ts_out=service=true:{OUT_DIR}',
           os.path.basename(f)
           ]
    print(' '.join(cmd))
    subprocess.run(cmd, shell=False, cwd='proto')