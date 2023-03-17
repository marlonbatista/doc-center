DROP DATABASE IF EXISTS DbCenter;
CREATE DATABASE DbCenter;
USE DbCenter;

CREATE TABLE Usuario (
	Id INT NOT NULL auto_increment,
	Nome VARCHAR(200) NULL,
	CPF VARCHAR(11) NULL,
    DataNascimento DATETIME NULL,
    CNPJ VARCHAR(14) NULL,
    PessoaFisica BIT,
    RazaoSocial VARCHAR(400) NULL,
    NomeFantasia VARCHAR(400) NULL,
    Email VARCHAR(300) NOT NULL,
	Senha TEXT NOT NULL,
	CONSTRAINT PK_Usuario PRIMARY KEY (Id),
    CONSTRAINT CK_USUARIO_PessoaFisica CHECK((LENGTH(CPF) > 0 AND PessoaFisica = 1) OR (LENGTH(CNPJ) > 0 AND PessoaFisica = 0))
);

CREATE TABLE Depedente (
	Id INT NOT NULL auto_increment,
    Nome VARCHAR(200) NOT NULL,
    DataNascimento DATETIME NOT NULL,
    IdResponsavel INT NOT NULL,
    CONSTRAINT PK_Dependente PRIMARY KEY (Id),
    CONSTRAINT FK_Dependente_Usuario FOREIGN KEY (IdResponsavel) REFERENCES Usuario (Id)
);

CREATE TABLE Arquivo (
	Id INT NOT NULL auto_increment,
	Descricao VARCHAR(200) NOT NULL,
	Arquivo BLOB(65535) NULL,
	IdUsuario INT NOT NULL,
	CONSTRAINT PK_Arquivo PRIMARY KEY (Id),
	CONSTRAINT FK_Arquivo_Usuario FOREIGN KEY (IdUsuario) REFERENCES Usuario(Id)
);

CREATE TABLE TipoPermissao(
	Id INT NOT NULL auto_increment,
	Descricao VARCHAR(200) NOT NULL,
	CONSTRAINT PK_TipoPermissao PRIMARY KEY (Id)
);

CREATE TABLE Permissao (
	Id INT NOT NULL auto_increment,
    IdTipoPermissao INT NOT NULL,
    Permanente BIT NOT NULL,
	IdArquivo INT NOT NULL,
	IdUsuarioDono INT NOT NULL,
	IdUsuarioAcesso INT NOT NULL,
	TempoDuracao DATETIME NULL,    
	CONSTRAINT PK_Permissao PRIMARY KEY (Id),
    constraint FK_Permissao_TipoPermissao FOREIGN KEY (IdTipoPermissao) REFERENCES TipoPermissao (Id),
	CONSTRAINT FK_Permissao_UsuarioDono FOREIGN KEY (IdUsuarioDono) REFERENCES Usuario (Id),
	CONSTRAINT FK_Permissao_UsuarioAcesso FOREIGN KEY (IdUsuarioAcesso) REFERENCES Usuario (Id),
    CONSTRAINT CK_Permissao_permanente CHECK(TempoDuracao IS NOT NULL AND Permanente = 0)
);