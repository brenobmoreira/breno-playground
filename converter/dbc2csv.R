# Set library paths
.libPaths(c("C:/R/library", .libPaths()))

# Load required library
if (!require(read.dbc, quietly = TRUE)) {
  stop("O pacote 'read.dbc' não está instalado. Por favor, instale-o com: install.packages('read.dbc')")
}

# Main function to convert DBC to CSV
dbc_to_csv <- function(file_path) {
  tryCatch({
    # Input validation
    if (missing(file_path) || is.null(file_path) || is.na(file_path) || file_path == "") {
      stop("Caminho do arquivo não fornecido")
    }
    
    if (!file.exists(file_path)) {
      stop(paste("Arquivo não encontrado:", file_path))
    }
    
    if (tools::file_ext(file_path) != "dbc") {
      stop("O arquivo fornecido não possui a extensão .dbc")
    }

    # Read DBC file
    dbc_data <- read.dbc(file_path)
    
    # Define output file path
    csv_file <- paste0(file_path, ".csv")
    
    # Write CSV file
    write.csv(dbc_data, csv_file, row.names = FALSE)
    
    # Verify file was created
    if (!file.exists(csv_file)) {
      stop("Falha ao criar o arquivo CSV de saída")
    }
    
    message("Conversão concluída com sucesso!")
    message("Arquivo de entrada: ", file_path)
    message("Arquivo de saída: ", csv_file)
    
    return(TRUE)
  }, error = function(e) {
    message("Erro na conversão: ", conditionMessage(e))
    return(FALSE)
  })
}

# Only run if executed directly (not sourced)
if (!interactive() && !is.null(commandArgs(trailingOnly = TRUE)[1])) {
  args <- commandArgs(trailingOnly = TRUE)
  quit(status = ifelse(dbc_to_csv(args[1]), 0, 1))
}
