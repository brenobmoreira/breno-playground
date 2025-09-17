import subprocess
import os
from pathlib import Path

def dbc_to_csv(file_path):

    try:
        script_dir = Path(__file__).parent
        r_script = script_dir / "dbc2csv.R"
        temp_dir = Path("R_temp")
        temp_dir.mkdir(exist_ok=True, parents=True)

        env = os.environ.copy()
        env['TEMP'] = str(temp_dir)
        env['TMP'] = str(temp_dir)
        env['TMPDIR'] = str(temp_dir)

        r_command = """
        source('{0}')
        """.format(
            r_script.resolve().as_posix()
        )

        temp_r_script = temp_dir / "temp_script.R"
        with open(temp_r_script, 'w', encoding='utf-8') as f:
            f.write(r_command)

        command = [
            "Rscript",
            "--vanilla",
            str(temp_r_script),
            str(Path(file_path).resolve())
        ]

        print(f"Executando comando: {' '.join(command)}")

        result = subprocess.run(
            command,
            capture_output=True,
            text=True,
            check=True,
            env=env,
            encoding='utf-8',
            errors='replace'
        )
        
        print("Script R executado com sucesso!")
        if result.stdout:
            print("--- SAÍDA (STDOUT) ---")
            print(result.stdout)
        if result.stderr:
            print("--- MENSAGENS (STDERR) ---")
            print(result.stderr)
        return True

    except subprocess.CalledProcessError as e:
        print("--- ERRO AO EXECUTAR O SCRIPT R ---")
        if e.stdout:
            print("STDOUT:", e.stdout)
        if e.stderr:
            print("STDERR:", e.stderr)
        return False
    finally:
        try:
            if 'temp_r_script' in locals() and temp_r_script.exists():
                temp_r_script.unlink()
        except Exception as e:
            print(f"Aviso: Não foi possível limpar o arquivo temporário: {e}")

dbc_to_csv("teste.dbc")