// Generated from c:\Users\email\Documents\GitHub\voltable\graph\query\openCypher\Cypher.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CypherLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		UNION=46, ALL=47, OPTIONAL=48, MATCH=49, UNWIND=50, AS=51, MERGE=52, ON=53, 
		CREATE=54, SET=55, DETACH=56, DELETE=57, REMOVE=58, CALL=59, YIELD=60, 
		WITH=61, DISTINCT=62, RETURN=63, ORDER=64, BY=65, L_SKIP=66, LIMIT=67, 
		ASCENDING=68, ASC=69, DESCENDING=70, DESC=71, WHERE=72, OR=73, XOR=74, 
		AND=75, NOT=76, IN=77, STARTS=78, ENDS=79, CONTAINS=80, IS=81, NULL=82, 
		COUNT=83, ANY=84, NONE=85, SINGLE=86, TRUE=87, FALSE=88, EXISTS=89, CASE=90, 
		ELSE=91, END=92, WHEN=93, THEN=94, StringLiteral=95, EscapedChar=96, HexInteger=97, 
		DecimalInteger=98, OctalInteger=99, HexLetter=100, HexDigit=101, Digit=102, 
		NonZeroDigit=103, NonZeroOctDigit=104, OctDigit=105, ZeroDigit=106, ExponentDecimalReal=107, 
		RegularDecimalReal=108, CONSTRAINT=109, DO=110, FOR=111, REQUIRE=112, 
		UNIQUE=113, MANDATORY=114, SCALAR=115, OF=116, ADD=117, DROP=118, FILTER=119, 
		EXTRACT=120, UnescapedSymbolicName=121, IdentifierStart=122, IdentifierPart=123, 
		EscapedSymbolicName=124, SP=125, WHITESPACE=126, Comment=127;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32", 
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40", 
		"T__41", "T__42", "T__43", "T__44", "UNION", "ALL", "OPTIONAL", "MATCH", 
		"UNWIND", "AS", "MERGE", "ON", "CREATE", "SET", "DETACH", "DELETE", "REMOVE", 
		"CALL", "YIELD", "WITH", "DISTINCT", "RETURN", "ORDER", "BY", "L_SKIP", 
		"LIMIT", "ASCENDING", "ASC", "DESCENDING", "DESC", "WHERE", "OR", "XOR", 
		"AND", "NOT", "IN", "STARTS", "ENDS", "CONTAINS", "IS", "NULL", "COUNT", 
		"ANY", "NONE", "SINGLE", "TRUE", "FALSE", "EXISTS", "CASE", "ELSE", "END", 
		"WHEN", "THEN", "StringLiteral", "EscapedChar", "HexInteger", "DecimalInteger", 
		"OctalInteger", "HexLetter", "HexDigit", "Digit", "NonZeroDigit", "NonZeroOctDigit", 
		"OctDigit", "ZeroDigit", "ExponentDecimalReal", "RegularDecimalReal", 
		"CONSTRAINT", "DO", "FOR", "REQUIRE", "UNIQUE", "MANDATORY", "SCALAR", 
		"OF", "ADD", "DROP", "FILTER", "EXTRACT", "UnescapedSymbolicName", "IdentifierStart", 
		"IdentifierPart", "EscapedSymbolicName", "SP", "WHITESPACE", "Comment", 
		"FF", "EscapedSymbolicName_0", "RS", "ID_Continue", "Comment_1", "StringLiteral_1", 
		"Comment_3", "Comment_2", "GS", "FS", "CR", "Sc", "SPACE", "Pc", "TAB", 
		"StringLiteral_0", "LF", "VT", "US", "ID_Start"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "';'", "','", "'='", "'+='", "'*'", "'('", "')'", "'['", "']'", 
		"':'", "'|'", "'..'", "'+'", "'-'", "'/'", "'%'", "'^'", "'<>'", "'<'", 
		"'>'", "'<='", "'>='", "'.'", "'{'", "'}'", "'$'", "'\u00E2\u0178\u00A8'", 
		"'\u00E3\u20AC\u02C6'", "'\u00EF\u00B9\u00A4'", "'\u00EF\u00BC\u0153'", 
		"'\u00E2\u0178\u00A9'", "'\u00E3\u20AC\u2030'", "'\u00EF\u00B9\u00A5'", 
		"'\u00EF\u00BC\u017E'", "'\u00C2\u00AD'", "'\u00E2\u20AC\uFFFD'", "'\u00E2\u20AC\u2018'", 
		"'\u00E2\u20AC\u2019'", "'\u00E2\u20AC\u201C'", "'\u00E2\u20AC\u201D'", 
		"'\u00E2\u20AC\u2022'", "'\u00E2\u02C6\u2019'", "'\u00EF\u00B9\u02DC'", 
		"'\u00EF\u00B9\u00A3'", "'\u00EF\u00BC\uFFFD'", null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, "'0'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, "UNION", "ALL", 
		"OPTIONAL", "MATCH", "UNWIND", "AS", "MERGE", "ON", "CREATE", "SET", "DETACH", 
		"DELETE", "REMOVE", "CALL", "YIELD", "WITH", "DISTINCT", "RETURN", "ORDER", 
		"BY", "L_SKIP", "LIMIT", "ASCENDING", "ASC", "DESCENDING", "DESC", "WHERE", 
		"OR", "XOR", "AND", "NOT", "IN", "STARTS", "ENDS", "CONTAINS", "IS", "NULL", 
		"COUNT", "ANY", "NONE", "SINGLE", "TRUE", "FALSE", "EXISTS", "CASE", "ELSE", 
		"END", "WHEN", "THEN", "StringLiteral", "EscapedChar", "HexInteger", "DecimalInteger", 
		"OctalInteger", "HexLetter", "HexDigit", "Digit", "NonZeroDigit", "NonZeroOctDigit", 
		"OctDigit", "ZeroDigit", "ExponentDecimalReal", "RegularDecimalReal", 
		"CONSTRAINT", "DO", "FOR", "REQUIRE", "UNIQUE", "MANDATORY", "SCALAR", 
		"OF", "ADD", "DROP", "FILTER", "EXTRACT", "UnescapedSymbolicName", "IdentifierStart", 
		"IdentifierPart", "EscapedSymbolicName", "SP", "WHITESPACE", "Comment"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public CypherLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Cypher.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\u0081\u043f\b\1\4"+
		"\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n"+
		"\4\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t"+
		"=\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4"+
		"I\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\t"+
		"T\4U\tU\4V\tV\4W\tW\4X\tX\4Y\tY\4Z\tZ\4[\t[\4\\\t\\\4]\t]\4^\t^\4_\t_"+
		"\4`\t`\4a\ta\4b\tb\4c\tc\4d\td\4e\te\4f\tf\4g\tg\4h\th\4i\ti\4j\tj\4k"+
		"\tk\4l\tl\4m\tm\4n\tn\4o\to\4p\tp\4q\tq\4r\tr\4s\ts\4t\tt\4u\tu\4v\tv"+
		"\4w\tw\4x\tx\4y\ty\4z\tz\4{\t{\4|\t|\4}\t}\4~\t~\4\177\t\177\4\u0080\t"+
		"\u0080\4\u0081\t\u0081\4\u0082\t\u0082\4\u0083\t\u0083\4\u0084\t\u0084"+
		"\4\u0085\t\u0085\4\u0086\t\u0086\4\u0087\t\u0087\4\u0088\t\u0088\4\u0089"+
		"\t\u0089\4\u008a\t\u008a\4\u008b\t\u008b\4\u008c\t\u008c\4\u008d\t\u008d"+
		"\4\u008e\t\u008e\4\u008f\t\u008f\4\u0090\t\u0090\4\u0091\t\u0091\4\u0092"+
		"\t\u0092\4\u0093\t\u0093\4\u0094\t\u0094\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3"+
		"\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3"+
		"\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\23"+
		"\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\31\3\31"+
		"\3\32\3\32\3\33\3\33\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\36\3\36"+
		"\3\36\3\36\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\""+
		"\3#\3#\3#\3#\3$\3$\3$\3%\3%\3%\3%\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3(\3(\3"+
		"(\3(\3)\3)\3)\3)\3*\3*\3*\3*\3+\3+\3+\3+\3,\3,\3,\3,\3-\3-\3-\3-\3.\3"+
		".\3.\3.\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61"+
		"\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\63"+
		"\3\63\3\63\3\63\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66"+
		"\3\66\3\67\3\67\3\67\3\67\3\67\3\67\3\67\38\38\38\38\39\39\39\39\39\3"+
		"9\39\3:\3:\3:\3:\3:\3:\3:\3;\3;\3;\3;\3;\3;\3;\3<\3<\3<\3<\3<\3=\3=\3"+
		"=\3=\3=\3=\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3?\3?\3?\3?\3@\3@\3@\3@\3@\3"+
		"@\3@\3A\3A\3A\3A\3A\3A\3B\3B\3B\3C\3C\3C\3C\3C\3D\3D\3D\3D\3D\3D\3E\3"+
		"E\3E\3E\3E\3E\3E\3E\3E\3E\3F\3F\3F\3F\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3"+
		"G\3H\3H\3H\3H\3H\3I\3I\3I\3I\3I\3I\3J\3J\3J\3K\3K\3K\3K\3L\3L\3L\3L\3"+
		"M\3M\3M\3M\3N\3N\3N\3O\3O\3O\3O\3O\3O\3O\3P\3P\3P\3P\3P\3Q\3Q\3Q\3Q\3"+
		"Q\3Q\3Q\3Q\3Q\3R\3R\3R\3S\3S\3S\3S\3S\3T\3T\3T\3T\3T\3T\3U\3U\3U\3U\3"+
		"V\3V\3V\3V\3V\3W\3W\3W\3W\3W\3W\3W\3X\3X\3X\3X\3X\3Y\3Y\3Y\3Y\3Y\3Y\3"+
		"Z\3Z\3Z\3Z\3Z\3Z\3Z\3[\3[\3[\3[\3[\3\\\3\\\3\\\3\\\3\\\3]\3]\3]\3]\3^"+
		"\3^\3^\3^\3^\3_\3_\3_\3_\3_\3`\3`\3`\7`\u02c4\n`\f`\16`\u02c7\13`\3`\3"+
		"`\3`\3`\7`\u02cd\n`\f`\16`\u02d0\13`\3`\5`\u02d3\n`\3a\3a\3a\3a\3a\3a"+
		"\3a\3a\3a\3a\3a\3a\3a\3a\3a\3a\3a\3a\5a\u02e7\na\3b\3b\3b\3b\6b\u02ed"+
		"\nb\rb\16b\u02ee\3c\3c\3c\7c\u02f4\nc\fc\16c\u02f7\13c\5c\u02f9\nc\3d"+
		"\3d\6d\u02fd\nd\rd\16d\u02fe\3e\5e\u0302\ne\3f\3f\5f\u0306\nf\3g\3g\5"+
		"g\u030a\ng\3h\3h\5h\u030e\nh\3i\3i\3j\3j\5j\u0314\nj\3k\3k\3l\6l\u0319"+
		"\nl\rl\16l\u031a\3l\6l\u031e\nl\rl\16l\u031f\3l\3l\6l\u0324\nl\rl\16l"+
		"\u0325\3l\3l\6l\u032a\nl\rl\16l\u032b\5l\u032e\nl\3l\5l\u0331\nl\3l\5"+
		"l\u0334\nl\3l\6l\u0337\nl\rl\16l\u0338\3m\7m\u033c\nm\fm\16m\u033f\13"+
		"m\3m\3m\6m\u0343\nm\rm\16m\u0344\3n\3n\3n\3n\3n\3n\3n\3n\3n\3n\3n\3o\3"+
		"o\3o\3p\3p\3p\3p\3q\3q\3q\3q\3q\3q\3q\3q\3r\3r\3r\3r\3r\3r\3r\3s\3s\3"+
		"s\3s\3s\3s\3s\3s\3s\3s\3t\3t\3t\3t\3t\3t\3t\3u\3u\3u\3v\3v\3v\3v\3w\3"+
		"w\3w\3w\3w\3x\3x\3x\3x\3x\3x\3x\3y\3y\3y\3y\3y\3y\3y\3y\3z\3z\7z\u0396"+
		"\nz\fz\16z\u0399\13z\3{\3{\5{\u039d\n{\3|\3|\5|\u03a1\n|\3}\3}\7}\u03a5"+
		"\n}\f}\16}\u03a8\13}\3}\6}\u03ab\n}\r}\16}\u03ac\3~\6~\u03b0\n~\r~\16"+
		"~\u03b1\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3"+
		"\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177"+
		"\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177"+
		"\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177"+
		"\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177"+
		"\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\3\177\5\177\u03f7\n\177"+
		"\3\u0080\3\u0080\3\u0080\3\u0080\3\u0080\3\u0080\7\u0080\u03ff\n\u0080"+
		"\f\u0080\16\u0080\u0402\13\u0080\3\u0080\3\u0080\3\u0080\3\u0080\3\u0080"+
		"\3\u0080\7\u0080\u040a\n\u0080\f\u0080\16\u0080\u040d\13\u0080\3\u0080"+
		"\5\u0080\u0410\n\u0080\3\u0080\3\u0080\5\u0080\u0414\n\u0080\5\u0080\u0416"+
		"\n\u0080\3\u0081\3\u0081\3\u0082\3\u0082\3\u0083\3\u0083\3\u0084\3\u0084"+
		"\3\u0085\3\u0085\3\u0086\3\u0086\3\u0087\3\u0087\3\u0088\3\u0088\3\u0089"+
		"\3\u0089\3\u008a\3\u008a\3\u008b\3\u008b\3\u008c\3\u008c\3\u008d\3\u008d"+
		"\3\u008e\3\u008e\3\u008f\3\u008f\3\u0090\3\u0090\3\u0091\3\u0091\3\u0092"+
		"\3\u0092\3\u0093\3\u0093\3\u0094\3\u0094\2\2\u0095\3\3\5\4\7\5\t\6\13"+
		"\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'"+
		"\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'"+
		"M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177"+
		"A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008fI\u0091J\u0093"+
		"K\u0095L\u0097M\u0099N\u009bO\u009dP\u009fQ\u00a1R\u00a3S\u00a5T\u00a7"+
		"U\u00a9V\u00abW\u00adX\u00afY\u00b1Z\u00b3[\u00b5\\\u00b7]\u00b9^\u00bb"+
		"_\u00bd`\u00bfa\u00c1b\u00c3c\u00c5d\u00c7e\u00c9f\u00cbg\u00cdh\u00cf"+
		"i\u00d1j\u00d3k\u00d5l\u00d7m\u00d9n\u00dbo\u00ddp\u00dfq\u00e1r\u00e3"+
		"s\u00e5t\u00e7u\u00e9v\u00ebw\u00edx\u00efy\u00f1z\u00f3{\u00f5|\u00f7"+
		"}\u00f9~\u00fb\177\u00fd\u0080\u00ff\u0081\u0101\2\u0103\2\u0105\2\u0107"+
		"\2\u0109\2\u010b\2\u010d\2\u010f\2\u0111\2\u0113\2\u0115\2\u0117\2\u0119"+
		"\2\u011b\2\u011d\2\u011f\2\u0121\2\u0123\2\u0125\2\u0127\2\3\2\60\4\2"+
		"WWww\4\2PPpp\4\2KKkk\4\2QQqq\4\2CCcc\4\2NNnn\4\2RRrr\4\2VVvv\4\2OOoo\4"+
		"\2EEee\4\2JJjj\4\2YYyy\4\2FFff\4\2UUuu\4\2GGgg\4\2TTtt\4\2IIii\4\2XXx"+
		"x\4\2[[{{\4\2DDdd\4\2MMmm\4\2ZZzz\4\2HHhh\17\2$$))DDHHPPTTVV^^ddhhppt"+
		"tvv\4\2CHch\4\2SSss\3\2\16\16\4\2\2ac\1\3\2  \u01af\2\62;C\\aac|\u00ac"+
		"\u00ac\u00b7\u00b7\u00b9\u00b9\u00bc\u00bc\u00c2\u00d8\u00da\u00f8\u00fa"+
		"\u02c3\u02c8\u02d3\u02e2\u02e6\u02ee\u02ee\u02f0\u02f0\u0302\u0376\u0378"+
		"\u0379\u037c\u037f\u0388\u038c\u038e\u038e\u0390\u03a3\u03a5\u03f7\u03f9"+
		"\u0483\u0485\u0489\u048c\u0529\u0533\u0558\u055b\u055b\u0563\u0589\u0593"+
		"\u05bf\u05c1\u05c1\u05c3\u05c4\u05c6\u05c7\u05c9\u05c9\u05d2\u05ec\u05f2"+
		"\u05f4\u0612\u061c\u0622\u066b\u0670\u06d5\u06d7\u06de\u06e1\u06ea\u06ec"+
		"\u06fe\u0701\u0701\u0712\u074c\u074f\u07b3\u07c2\u07f7\u07fc\u07fc\u0802"+
		"\u082f\u0842\u085d\u08a2\u08a2\u08a4\u08ae\u08e6\u0900\u0902\u0965\u0968"+
		"\u0971\u0973\u0979\u097b\u0981\u0983\u0985\u0987\u098e\u0991\u0992\u0995"+
		"\u09aa\u09ac\u09b2\u09b4\u09b4\u09b8\u09bb\u09be\u09c6\u09c9\u09ca\u09cd"+
		"\u09d0\u09d9\u09d9\u09de\u09df\u09e1\u09e5\u09e8\u09f3\u0a03\u0a05\u0a07"+
		"\u0a0c\u0a11\u0a12\u0a15\u0a2a\u0a2c\u0a32\u0a34\u0a35\u0a37\u0a38\u0a3a"+
		"\u0a3b\u0a3e\u0a3e\u0a40\u0a44\u0a49\u0a4a\u0a4d\u0a4f\u0a53\u0a53\u0a5b"+
		"\u0a5e\u0a60\u0a60\u0a68\u0a77\u0a83\u0a85\u0a87\u0a8f\u0a91\u0a93\u0a95"+
		"\u0aaa\u0aac\u0ab2\u0ab4\u0ab5\u0ab7\u0abb\u0abe\u0ac7\u0ac9\u0acb\u0acd"+
		"\u0acf\u0ad2\u0ad2\u0ae2\u0ae5\u0ae8\u0af1\u0b03\u0b05\u0b07\u0b0e\u0b11"+
		"\u0b12\u0b15\u0b2a\u0b2c\u0b32\u0b34\u0b35\u0b37\u0b3b\u0b3e\u0b46\u0b49"+
		"\u0b4a\u0b4d\u0b4f\u0b58\u0b59\u0b5e\u0b5f\u0b61\u0b65\u0b68\u0b71\u0b73"+
		"\u0b73\u0b84\u0b85\u0b87\u0b8c\u0b90\u0b92\u0b94\u0b97\u0b9b\u0b9c\u0b9e"+
		"\u0b9e\u0ba0\u0ba1\u0ba5\u0ba6\u0baa\u0bac\u0bb0\u0bbb\u0bc0\u0bc4\u0bc8"+
		"\u0bca\u0bcc\u0bcf\u0bd2\u0bd2\u0bd9\u0bd9\u0be8\u0bf1\u0c03\u0c05\u0c07"+
		"\u0c0e\u0c10\u0c12\u0c14\u0c2a\u0c2c\u0c35\u0c37\u0c3b\u0c3f\u0c46\u0c48"+
		"\u0c4a\u0c4c\u0c4f\u0c57\u0c58\u0c5a\u0c5b\u0c62\u0c65\u0c68\u0c71\u0c84"+
		"\u0c85\u0c87\u0c8e\u0c90\u0c92\u0c94\u0caa\u0cac\u0cb5\u0cb7\u0cbb\u0cbe"+
		"\u0cc6\u0cc8\u0cca\u0ccc\u0ccf\u0cd7\u0cd8\u0ce0\u0ce0\u0ce2\u0ce5\u0ce8"+
		"\u0cf1\u0cf3\u0cf4\u0d04\u0d05\u0d07\u0d0e\u0d10\u0d12\u0d14\u0d3c\u0d3f"+
		"\u0d46\u0d48\u0d4a\u0d4c\u0d50\u0d59\u0d59\u0d62\u0d65\u0d68\u0d71\u0d7c"+
		"\u0d81\u0d84\u0d85\u0d87\u0d98\u0d9c\u0db3\u0db5\u0dbd\u0dbf\u0dbf\u0dc2"+
		"\u0dc8\u0dcc\u0dcc\u0dd1\u0dd6\u0dd8\u0dd8\u0dda\u0de1\u0df4\u0df5\u0e03"+
		"\u0e3c\u0e42\u0e50\u0e52\u0e5b\u0e83\u0e84\u0e86\u0e86\u0e89\u0e8a\u0e8c"+
		"\u0e8c\u0e8f\u0e8f\u0e96\u0e99\u0e9b\u0ea1\u0ea3\u0ea5\u0ea7\u0ea7\u0ea9"+
		"\u0ea9\u0eac\u0ead\u0eaf\u0ebb\u0ebd\u0ebf\u0ec2\u0ec6\u0ec8\u0ec8\u0eca"+
		"\u0ecf\u0ed2\u0edb\u0ede\u0ee1\u0f02\u0f02\u0f1a\u0f1b\u0f22\u0f2b\u0f37"+
		"\u0f37\u0f39\u0f39\u0f3b\u0f3b\u0f40\u0f49\u0f4b\u0f6e\u0f73\u0f86\u0f88"+
		"\u0f99\u0f9b\u0fbe\u0fc8\u0fc8\u1002\u104b\u1052\u109f\u10a2\u10c7\u10c9"+
		"\u10c9\u10cf\u10cf\u10d2\u10fc\u10fe\u124a\u124c\u124f\u1252\u1258\u125a"+
		"\u125a\u125c\u125f\u1262\u128a\u128c\u128f\u1292\u12b2\u12b4\u12b7\u12ba"+
		"\u12c0\u12c2\u12c2\u12c4\u12c7\u12ca\u12d8\u12da\u1312\u1314\u1317\u131a"+
		"\u135c\u135f\u1361\u136b\u1373\u1382\u1391\u13a2\u13f6\u1403\u166e\u1671"+
		"\u1681\u1683\u169c\u16a2\u16ec\u16f0\u16f2\u1702\u170e\u1710\u1716\u1722"+
		"\u1736\u1742\u1755\u1762\u176e\u1770\u1772\u1774\u1775\u1782\u17d5\u17d9"+
		"\u17d9\u17de\u17df\u17e2\u17eb\u180d\u180f\u1812\u181b\u1822\u1879\u1882"+
		"\u18ac\u18b2\u18f7\u1902\u191e\u1922\u192d\u1932\u193d\u1948\u196f\u1972"+
		"\u1976\u1982\u19ad\u19b2\u19cb\u19d2\u19dc\u1a02\u1a1d\u1a22\u1a60\u1a62"+
		"\u1a7e\u1a81\u1a8b\u1a92\u1a9b\u1aa9\u1aa9\u1b02\u1b4d\u1b52\u1b5b\u1b6d"+
		"\u1b75\u1b82\u1bf5\u1c02\u1c39\u1c42\u1c4b\u1c4f\u1c7f\u1cd2\u1cd4\u1cd6"+
		"\u1cf8\u1d02\u1de8\u1dfe\u1f17\u1f1a\u1f1f\u1f22\u1f47\u1f4a\u1f4f\u1f52"+
		"\u1f59\u1f5b\u1f5b\u1f5d\u1f5d\u1f5f\u1f5f\u1f61\u1f7f\u1f82\u1fb6\u1fb8"+
		"\u1fbe\u1fc0\u1fc0\u1fc4\u1fc6\u1fc8\u1fce\u1fd2\u1fd5\u1fd8\u1fdd\u1fe2"+
		"\u1fee\u1ff4\u1ff6\u1ff8\u1ffe\u2041\u2042\u2056\u2056\u2073\u2073\u2081"+
		"\u2081\u2092\u209e\u20d2\u20de\u20e3\u20e3\u20e7\u20f2\u2104\u2104\u2109"+
		"\u2109\u210c\u2115\u2117\u2117\u211a\u211f\u2126\u2126\u2128\u2128\u212a"+
		"\u212a\u212c\u213b\u213e\u2141\u2147\u214b\u2150\u2150\u2162\u218a\u2c02"+
		"\u2c30\u2c32\u2c60\u2c62\u2ce6\u2ced\u2cf5\u2d02\u2d27\u2d29\u2d29\u2d2f"+
		"\u2d2f\u2d32\u2d69\u2d71\u2d71\u2d81\u2d98\u2da2\u2da8\u2daa\u2db0\u2db2"+
		"\u2db8\u2dba\u2dc0\u2dc2\u2dc8\u2dca\u2dd0\u2dd2\u2dd8\u2dda\u2de0\u2de2"+
		"\u2e01\u3007\u3009\u3023\u3031\u3033\u3037\u303a\u303e\u3043\u3098\u309b"+
		"\u30a1\u30a3\u30fc\u30fe\u3101\u3107\u312f\u3133\u3190\u31a2\u31bc\u31f2"+
		"\u3201\u3402\u4db7\u4e02\u9fce\ua002\ua48e\ua4d2\ua4ff\ua502\ua60e\ua612"+
		"\ua62d\ua642\ua671\ua676\ua67f\ua681\ua699\ua6a1\ua6f3\ua719\ua721\ua724"+
		"\ua78a\ua78d\ua790\ua792\ua795\ua7a2\ua7ac\ua7fa\ua829\ua842\ua875\ua882"+
		"\ua8c6\ua8d2\ua8db\ua8e2\ua8f9\ua8fd\ua8fd\ua902\ua92f\ua932\ua955\ua962"+
		"\ua97e\ua982\ua9c2\ua9d1\ua9db\uaa02\uaa38\uaa42\uaa4f\uaa52\uaa5b\uaa62"+
		"\uaa78\uaa7c\uaa7d\uaa82\uaac4\uaadd\uaadf\uaae2\uaaf1\uaaf4\uaaf8\uab03"+
		"\uab08\uab0b\uab10\uab13\uab18\uab22\uab28\uab2a\uab30\uabc2\uabec\uabee"+
		"\uabef\uabf2\uabfb\uac02\ud7a5\ud7b2\ud7c8\ud7cd\ud7fd\uf902\ufa6f\ufa72"+
		"\ufadb\ufb02\ufb08\ufb15\ufb19\ufb1f\ufb2a\ufb2c\ufb38\ufb3a\ufb3e\ufb40"+
		"\ufb40\ufb42\ufb43\ufb45\ufb46\ufb48\ufbb3\ufbd5\ufd3f\ufd52\ufd91\ufd94"+
		"\ufdc9\ufdf2\ufdfd\ufe02\ufe11\ufe22\ufe28\ufe35\ufe36\ufe4f\ufe51\ufe72"+
		"\ufe76\ufe78\ufefe\uff12\uff1b\uff23\uff3c\uff41\uff41\uff43\uff5c\uff68"+
		"\uffc0\uffc4\uffc9\uffcc\uffd1\uffd4\uffd9\uffdc\uffde\4\2\2+-\1\5\2\2"+
		"(*]_\1\5\2\2\13\r\16\20\1\4\2\2\60\62\1\3\2\37\37\3\2\36\36\3\2\17\17"+
		"\23\2&&\u00a4\u00a7\u0591\u0591\u060d\u060d\u09f4\u09f5\u09fd\u09fd\u0af3"+
		"\u0af3\u0bfb\u0bfb\u0e41\u0e41\u17dd\u17dd\u20a2\u20bc\ua83a\ua83a\ufdfe"+
		"\ufdfe\ufe6b\ufe6b\uff06\uff06\uffe2\uffe3\uffe7\uffe8\3\2\"\"\b\2aa\u2041"+
		"\u2042\u2056\u2056\ufe35\ufe36\ufe4f\ufe51\uff41\uff41\3\2\13\13\5\2\2"+
		"#%]_\1\3\2\f\f\3\2\r\r\3\2!!\u0174\2C\\c|\u00ac\u00ac\u00b7\u00b7\u00bc"+
		"\u00bc\u00c2\u00d8\u00da\u00f8\u00fa\u02c3\u02c8\u02d3\u02e2\u02e6\u02ee"+
		"\u02ee\u02f0\u02f0\u0372\u0376\u0378\u0379\u037c\u037f\u0388\u0388\u038a"+
		"\u038c\u038e\u038e\u0390\u03a3\u03a5\u03f7\u03f9\u0483\u048c\u0529\u0533"+
		"\u0558\u055b\u055b\u0563\u0589\u05d2\u05ec\u05f2\u05f4\u0622\u064c\u0670"+
		"\u0671\u0673\u06d5\u06d7\u06d7\u06e7\u06e8\u06f0\u06f1\u06fc\u06fe\u0701"+
		"\u0701\u0712\u0712\u0714\u0731\u074f\u07a7\u07b3\u07b3\u07cc\u07ec\u07f6"+
		"\u07f7\u07fc\u07fc\u0802\u0817\u081c\u081c\u0826\u0826\u082a\u082a\u0842"+
		"\u085a\u08a2\u08a2\u08a4\u08ae\u0906\u093b\u093f\u093f\u0952\u0952\u095a"+
		"\u0963\u0973\u0979\u097b\u0981\u0987\u098e\u0991\u0992\u0995\u09aa\u09ac"+
		"\u09b2\u09b4\u09b4\u09b8\u09bb\u09bf\u09bf\u09d0\u09d0\u09de\u09df\u09e1"+
		"\u09e3\u09f2\u09f3\u0a07\u0a0c\u0a11\u0a12\u0a15\u0a2a\u0a2c\u0a32\u0a34"+
		"\u0a35\u0a37\u0a38\u0a3a\u0a3b\u0a5b\u0a5e\u0a60\u0a60\u0a74\u0a76\u0a87"+
		"\u0a8f\u0a91\u0a93\u0a95\u0aaa\u0aac\u0ab2\u0ab4\u0ab5\u0ab7\u0abb\u0abf"+
		"\u0abf\u0ad2\u0ad2\u0ae2\u0ae3\u0b07\u0b0e\u0b11\u0b12\u0b15\u0b2a\u0b2c"+
		"\u0b32\u0b34\u0b35\u0b37\u0b3b\u0b3f\u0b3f\u0b5e\u0b5f\u0b61\u0b63\u0b73"+
		"\u0b73\u0b85\u0b85\u0b87\u0b8c\u0b90\u0b92\u0b94\u0b97\u0b9b\u0b9c\u0b9e"+
		"\u0b9e\u0ba0\u0ba1\u0ba5\u0ba6\u0baa\u0bac\u0bb0\u0bbb\u0bd2\u0bd2\u0c07"+
		"\u0c0e\u0c10\u0c12\u0c14\u0c2a\u0c2c\u0c35\u0c37\u0c3b\u0c3f\u0c3f\u0c5a"+
		"\u0c5b\u0c62\u0c63\u0c87\u0c8e\u0c90\u0c92\u0c94\u0caa\u0cac\u0cb5\u0cb7"+
		"\u0cbb\u0cbf\u0cbf\u0ce0\u0ce0\u0ce2\u0ce3\u0cf3\u0cf4\u0d07\u0d0e\u0d10"+
		"\u0d12\u0d14\u0d3c\u0d3f\u0d3f\u0d50\u0d50\u0d62\u0d63\u0d7c\u0d81\u0d87"+
		"\u0d98\u0d9c\u0db3\u0db5\u0dbd\u0dbf\u0dbf\u0dc2\u0dc8\u0e03\u0e32\u0e34"+
		"\u0e35\u0e42\u0e48\u0e83\u0e84\u0e86\u0e86\u0e89\u0e8a\u0e8c\u0e8c\u0e8f"+
		"\u0e8f\u0e96\u0e99\u0e9b\u0ea1\u0ea3\u0ea5\u0ea7\u0ea7\u0ea9\u0ea9\u0eac"+
		"\u0ead\u0eaf\u0eb2\u0eb4\u0eb5\u0ebf\u0ebf\u0ec2\u0ec6\u0ec8\u0ec8\u0ede"+
		"\u0ee1\u0f02\u0f02\u0f42\u0f49\u0f4b\u0f6e\u0f8a\u0f8e\u1002\u102c\u1041"+
		"\u1041\u1052\u1057\u105c\u105f\u1063\u1063\u1067\u1068\u1070\u1072\u1077"+
		"\u1083\u1090\u1090\u10a2\u10c7\u10c9\u10c9\u10cf\u10cf\u10d2\u10fc\u10fe"+
		"\u124a\u124c\u124f\u1252\u1258\u125a\u125a\u125c\u125f\u1262\u128a\u128c"+
		"\u128f\u1292\u12b2\u12b4\u12b7\u12ba\u12c0\u12c2\u12c2\u12c4\u12c7\u12ca"+
		"\u12d8\u12da\u1312\u1314\u1317\u131a\u135c\u1382\u1391\u13a2\u13f6\u1403"+
		"\u166e\u1671\u1681\u1683\u169c\u16a2\u16ec\u16f0\u16f2\u1702\u170e\u1710"+
		"\u1713\u1722\u1733\u1742\u1753\u1762\u176e\u1770\u1772\u1782\u17b5\u17d9"+
		"\u17d9\u17de\u17de\u1822\u1879\u1882\u18aa\u18ac\u18ac\u18b2\u18f7\u1902"+
		"\u191e\u1952\u196f\u1972\u1976\u1982\u19ad\u19c3\u19c9\u1a02\u1a18\u1a22"+
		"\u1a56\u1aa9\u1aa9\u1b07\u1b35\u1b47\u1b4d\u1b85\u1ba2\u1bb0\u1bb1\u1bbc"+
		"\u1be7\u1c02\u1c25\u1c4f\u1c51\u1c5c\u1c7f\u1ceb\u1cee\u1cf0\u1cf3\u1cf7"+
		"\u1cf8\u1d02\u1dc1\u1e02\u1f17\u1f1a\u1f1f\u1f22\u1f47\u1f4a\u1f4f\u1f52"+
		"\u1f59\u1f5b\u1f5b\u1f5d\u1f5d\u1f5f\u1f5f\u1f61\u1f7f\u1f82\u1fb6\u1fb8"+
		"\u1fbe\u1fc0\u1fc0\u1fc4\u1fc6\u1fc8\u1fce\u1fd2\u1fd5\u1fd8\u1fdd\u1fe2"+
		"\u1fee\u1ff4\u1ff6\u1ff8\u1ffe\u2073\u2073\u2081\u2081\u2092\u209e\u2104"+
		"\u2104\u2109\u2109\u210c\u2115\u2117\u2117\u211a\u211f\u2126\u2126\u2128"+
		"\u2128\u212a\u212a\u212c\u213b\u213e\u2141\u2147\u214b\u2150\u2150\u2162"+
		"\u218a\u2c02\u2c30\u2c32\u2c60\u2c62\u2ce6\u2ced\u2cf0\u2cf4\u2cf5\u2d02"+
		"\u2d27\u2d29\u2d29\u2d2f\u2d2f\u2d32\u2d69\u2d71\u2d71\u2d82\u2d98\u2da2"+
		"\u2da8\u2daa\u2db0\u2db2\u2db8\u2dba\u2dc0\u2dc2\u2dc8\u2dca\u2dd0\u2dd2"+
		"\u2dd8\u2dda\u2de0\u3007\u3009\u3023\u302b\u3033\u3037\u303a\u303e\u3043"+
		"\u3098\u309d\u30a1\u30a3\u30fc\u30fe\u3101\u3107\u312f\u3133\u3190\u31a2"+
		"\u31bc\u31f2\u3201\u3402\u4db7\u4e02\u9fce\ua002\ua48e\ua4d2\ua4ff\ua502"+
		"\ua60e\ua612\ua621\ua62c\ua62d\ua642\ua670\ua681\ua699\ua6a2\ua6f1\ua719"+
		"\ua721\ua724\ua78a\ua78d\ua790\ua792\ua795\ua7a2\ua7ac\ua7fa\ua803\ua805"+
		"\ua807\ua809\ua80c\ua80e\ua824\ua842\ua875\ua884\ua8b5\ua8f4\ua8f9\ua8fd"+
		"\ua8fd\ua90c\ua927\ua932\ua948\ua962\ua97e\ua986\ua9b4\ua9d1\ua9d1\uaa02"+
		"\uaa2a\uaa42\uaa44\uaa46\uaa4d\uaa62\uaa78\uaa7c\uaa7c\uaa82\uaab1\uaab3"+
		"\uaab3\uaab7\uaab8\uaabb\uaabf\uaac2\uaac2\uaac4\uaac4\uaadd\uaadf\uaae2"+
		"\uaaec\uaaf4\uaaf6\uab03\uab08\uab0b\uab10\uab13\uab18\uab22\uab28\uab2a"+
		"\uab30\uabc2\uabe4\uac02\ud7a5\ud7b2\ud7c8\ud7cd\ud7fd\uf902\ufa6f\ufa72"+
		"\ufadb\ufb02\ufb08\ufb15\ufb19\ufb1f\ufb1f\ufb21\ufb2a\ufb2c\ufb38\ufb3a"+
		"\ufb3e\ufb40\ufb40\ufb42\ufb43\ufb45\ufb46\ufb48\ufbb3\ufbd5\ufd3f\ufd52"+
		"\ufd91\ufd94\ufdc9\ufdf2\ufdfd\ufe72\ufe76\ufe78\ufefe\uff23\uff3c\uff43"+
		"\uff5c\uff68\uffc0\uffc4\uffc9\uffcc\uffd1\uffd4\uffd9\uffdc\uffde\2\u046c"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2"+
		"\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2"+
		"{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085"+
		"\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2"+
		"\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2\2\u0095\3\2\2\2\2\u0097"+
		"\3\2\2\2\2\u0099\3\2\2\2\2\u009b\3\2\2\2\2\u009d\3\2\2\2\2\u009f\3\2\2"+
		"\2\2\u00a1\3\2\2\2\2\u00a3\3\2\2\2\2\u00a5\3\2\2\2\2\u00a7\3\2\2\2\2\u00a9"+
		"\3\2\2\2\2\u00ab\3\2\2\2\2\u00ad\3\2\2\2\2\u00af\3\2\2\2\2\u00b1\3\2\2"+
		"\2\2\u00b3\3\2\2\2\2\u00b5\3\2\2\2\2\u00b7\3\2\2\2\2\u00b9\3\2\2\2\2\u00bb"+
		"\3\2\2\2\2\u00bd\3\2\2\2\2\u00bf\3\2\2\2\2\u00c1\3\2\2\2\2\u00c3\3\2\2"+
		"\2\2\u00c5\3\2\2\2\2\u00c7\3\2\2\2\2\u00c9\3\2\2\2\2\u00cb\3\2\2\2\2\u00cd"+
		"\3\2\2\2\2\u00cf\3\2\2\2\2\u00d1\3\2\2\2\2\u00d3\3\2\2\2\2\u00d5\3\2\2"+
		"\2\2\u00d7\3\2\2\2\2\u00d9\3\2\2\2\2\u00db\3\2\2\2\2\u00dd\3\2\2\2\2\u00df"+
		"\3\2\2\2\2\u00e1\3\2\2\2\2\u00e3\3\2\2\2\2\u00e5\3\2\2\2\2\u00e7\3\2\2"+
		"\2\2\u00e9\3\2\2\2\2\u00eb\3\2\2\2\2\u00ed\3\2\2\2\2\u00ef\3\2\2\2\2\u00f1"+
		"\3\2\2\2\2\u00f3\3\2\2\2\2\u00f5\3\2\2\2\2\u00f7\3\2\2\2\2\u00f9\3\2\2"+
		"\2\2\u00fb\3\2\2\2\2\u00fd\3\2\2\2\2\u00ff\3\2\2\2\3\u0129\3\2\2\2\5\u012b"+
		"\3\2\2\2\7\u012d\3\2\2\2\t\u012f\3\2\2\2\13\u0132\3\2\2\2\r\u0134\3\2"+
		"\2\2\17\u0136\3\2\2\2\21\u0138\3\2\2\2\23\u013a\3\2\2\2\25\u013c\3\2\2"+
		"\2\27\u013e\3\2\2\2\31\u0140\3\2\2\2\33\u0143\3\2\2\2\35\u0145\3\2\2\2"+
		"\37\u0147\3\2\2\2!\u0149\3\2\2\2#\u014b\3\2\2\2%\u014d\3\2\2\2\'\u0150"+
		"\3\2\2\2)\u0152\3\2\2\2+\u0154\3\2\2\2-\u0157\3\2\2\2/\u015a\3\2\2\2\61"+
		"\u015c\3\2\2\2\63\u015e\3\2\2\2\65\u0160\3\2\2\2\67\u0162\3\2\2\29\u0166"+
		"\3\2\2\2;\u016a\3\2\2\2=\u016e\3\2\2\2?\u0172\3\2\2\2A\u0176\3\2\2\2C"+
		"\u017a\3\2\2\2E\u017e\3\2\2\2G\u0182\3\2\2\2I\u0185\3\2\2\2K\u0189\3\2"+
		"\2\2M\u018d\3\2\2\2O\u0191\3\2\2\2Q\u0195\3\2\2\2S\u0199\3\2\2\2U\u019d"+
		"\3\2\2\2W\u01a1\3\2\2\2Y\u01a5\3\2\2\2[\u01a9\3\2\2\2]\u01ad\3\2\2\2_"+
		"\u01b3\3\2\2\2a\u01b7\3\2\2\2c\u01c0\3\2\2\2e\u01c6\3\2\2\2g\u01cd\3\2"+
		"\2\2i\u01d0\3\2\2\2k\u01d6\3\2\2\2m\u01d9\3\2\2\2o\u01e0\3\2\2\2q\u01e4"+
		"\3\2\2\2s\u01eb\3\2\2\2u\u01f2\3\2\2\2w\u01f9\3\2\2\2y\u01fe\3\2\2\2{"+
		"\u0204\3\2\2\2}\u0209\3\2\2\2\177\u0212\3\2\2\2\u0081\u0219\3\2\2\2\u0083"+
		"\u021f\3\2\2\2\u0085\u0222\3\2\2\2\u0087\u0227\3\2\2\2\u0089\u022d\3\2"+
		"\2\2\u008b\u0237\3\2\2\2\u008d\u023b\3\2\2\2\u008f\u0246\3\2\2\2\u0091"+
		"\u024b\3\2\2\2\u0093\u0251\3\2\2\2\u0095\u0254\3\2\2\2\u0097\u0258\3\2"+
		"\2\2\u0099\u025c\3\2\2\2\u009b\u0260\3\2\2\2\u009d\u0263\3\2\2\2\u009f"+
		"\u026a\3\2\2\2\u00a1\u026f\3\2\2\2\u00a3\u0278\3\2\2\2\u00a5\u027b\3\2"+
		"\2\2\u00a7\u0280\3\2\2\2\u00a9\u0286\3\2\2\2\u00ab\u028a\3\2\2\2\u00ad"+
		"\u028f\3\2\2\2\u00af\u0296\3\2\2\2\u00b1\u029b\3\2\2\2\u00b3\u02a1\3\2"+
		"\2\2\u00b5\u02a8\3\2\2\2\u00b7\u02ad\3\2\2\2\u00b9\u02b2\3\2\2\2\u00bb"+
		"\u02b6\3\2\2\2\u00bd\u02bb\3\2\2\2\u00bf\u02d2\3\2\2\2\u00c1\u02d4\3\2"+
		"\2\2\u00c3\u02e8\3\2\2\2\u00c5\u02f8\3\2\2\2\u00c7\u02fa\3\2\2\2\u00c9"+
		"\u0301\3\2\2\2\u00cb\u0305\3\2\2\2\u00cd\u0309\3\2\2\2\u00cf\u030d\3\2"+
		"\2\2\u00d1\u030f\3\2\2\2\u00d3\u0313\3\2\2\2\u00d5\u0315\3\2\2\2\u00d7"+
		"\u032d\3\2\2\2\u00d9\u033d\3\2\2\2\u00db\u0346\3\2\2\2\u00dd\u0351\3\2"+
		"\2\2\u00df\u0354\3\2\2\2\u00e1\u0358\3\2\2\2\u00e3\u0360\3\2\2\2\u00e5"+
		"\u0367\3\2\2\2\u00e7\u0371\3\2\2\2\u00e9\u0378\3\2\2\2\u00eb\u037b\3\2"+
		"\2\2\u00ed\u037f\3\2\2\2\u00ef\u0384\3\2\2\2\u00f1\u038b\3\2\2\2\u00f3"+
		"\u0393\3\2\2\2\u00f5\u039c\3\2\2\2\u00f7\u03a0\3\2\2\2\u00f9\u03aa\3\2"+
		"\2\2\u00fb\u03af\3\2\2\2\u00fd\u03f6\3\2\2\2\u00ff\u0415\3\2\2\2\u0101"+
		"\u0417\3\2\2\2\u0103\u0419\3\2\2\2\u0105\u041b\3\2\2\2\u0107\u041d\3\2"+
		"\2\2\u0109\u041f\3\2\2\2\u010b\u0421\3\2\2\2\u010d\u0423\3\2\2\2\u010f"+
		"\u0425\3\2\2\2\u0111\u0427\3\2\2\2\u0113\u0429\3\2\2\2\u0115\u042b\3\2"+
		"\2\2\u0117\u042d\3\2\2\2\u0119\u042f\3\2\2\2\u011b\u0431\3\2\2\2\u011d"+
		"\u0433\3\2\2\2\u011f\u0435\3\2\2\2\u0121\u0437\3\2\2\2\u0123\u0439\3\2"+
		"\2\2\u0125\u043b\3\2\2\2\u0127\u043d\3\2\2\2\u0129\u012a\7=\2\2\u012a"+
		"\4\3\2\2\2\u012b\u012c\7.\2\2\u012c\6\3\2\2\2\u012d\u012e\7?\2\2\u012e"+
		"\b\3\2\2\2\u012f\u0130\7-\2\2\u0130\u0131\7?\2\2\u0131\n\3\2\2\2\u0132"+
		"\u0133\7,\2\2\u0133\f\3\2\2\2\u0134\u0135\7*\2\2\u0135\16\3\2\2\2\u0136"+
		"\u0137\7+\2\2\u0137\20\3\2\2\2\u0138\u0139\7]\2\2\u0139\22\3\2\2\2\u013a"+
		"\u013b\7_\2\2\u013b\24\3\2\2\2\u013c\u013d\7<\2\2\u013d\26\3\2\2\2\u013e"+
		"\u013f\7~\2\2\u013f\30\3\2\2\2\u0140\u0141\7\60\2\2\u0141\u0142\7\60\2"+
		"\2\u0142\32\3\2\2\2\u0143\u0144\7-\2\2\u0144\34\3\2\2\2\u0145\u0146\7"+
		"/\2\2\u0146\36\3\2\2\2\u0147\u0148\7\61\2\2\u0148 \3\2\2\2\u0149\u014a"+
		"\7\'\2\2\u014a\"\3\2\2\2\u014b\u014c\7`\2\2\u014c$\3\2\2\2\u014d\u014e"+
		"\7>\2\2\u014e\u014f\7@\2\2\u014f&\3\2\2\2\u0150\u0151\7>\2\2\u0151(\3"+
		"\2\2\2\u0152\u0153\7@\2\2\u0153*\3\2\2\2\u0154\u0155\7>\2\2\u0155\u0156"+
		"\7?\2\2\u0156,\3\2\2\2\u0157\u0158\7@\2\2\u0158\u0159\7?\2\2\u0159.\3"+
		"\2\2\2\u015a\u015b\7\60\2\2\u015b\60\3\2\2\2\u015c\u015d\7}\2\2\u015d"+
		"\62\3\2\2\2\u015e\u015f\7\177\2\2\u015f\64\3\2\2\2\u0160\u0161\7&\2\2"+
		"\u0161\66\3\2\2\2\u0162\u0163\7\u00e4\2\2\u0163\u0164\7\u017a\2\2\u0164"+
		"\u0165\7\u00aa\2\2\u01658\3\2\2\2\u0166\u0167\7\u00e5\2\2\u0167\u0168"+
		"\7\u20ae\2\2\u0168\u0169\7\u02c8\2\2\u0169:\3\2\2\2\u016a\u016b\7\u00f1"+
		"\2\2\u016b\u016c\7\u00bb\2\2\u016c\u016d\7\u00a6\2\2\u016d<\3\2\2\2\u016e"+
		"\u016f\7\u00f1\2\2\u016f\u0170\7\u00be\2\2\u0170\u0171\7\u0155\2\2\u0171"+
		">\3\2\2\2\u0172\u0173\7\u00e4\2\2\u0173\u0174\7\u017a\2\2\u0174\u0175"+
		"\7\u00ab\2\2\u0175@\3\2\2\2\u0176\u0177\7\u00e5\2\2\u0177\u0178\7\u20ae"+
		"\2\2\u0178\u0179\7\u2032\2\2\u0179B\3\2\2\2\u017a\u017b\7\u00f1\2\2\u017b"+
		"\u017c\7\u00bb\2\2\u017c\u017d\7\u00a7\2\2\u017dD\3\2\2\2\u017e\u017f"+
		"\7\u00f1\2\2\u017f\u0180\7\u00be\2\2\u0180\u0181\7\u0180\2\2\u0181F\3"+
		"\2\2\2\u0182\u0183\7\u00c4\2\2\u0183\u0184\7\u00af\2\2\u0184H\3\2\2\2"+
		"\u0185\u0186\7\u00e4\2\2\u0186\u0187\7\u20ae\2\2\u0187\u0188\7\uffff\2"+
		"\2\u0188J\3\2\2\2\u0189\u018a\7\u00e4\2\2\u018a\u018b\7\u20ae\2\2\u018b"+
		"\u018c\7\u201a\2\2\u018cL\3\2\2\2\u018d\u018e\7\u00e4\2\2\u018e\u018f"+
		"\7\u20ae\2\2\u018f\u0190\7\u201b\2\2\u0190N\3\2\2\2\u0191\u0192\7\u00e4"+
		"\2\2\u0192\u0193\7\u20ae\2\2\u0193\u0194\7\u201e\2\2\u0194P\3\2\2\2\u0195"+
		"\u0196\7\u00e4\2\2\u0196\u0197\7\u20ae\2\2\u0197\u0198\7\u201f\2\2\u0198"+
		"R\3\2\2\2\u0199\u019a\7\u00e4\2\2\u019a\u019b\7\u20ae\2\2\u019b\u019c"+
		"\7\u2024\2\2\u019cT\3\2\2\2\u019d\u019e\7\u00e4\2\2\u019e\u019f\7\u02c8"+
		"\2\2\u019f\u01a0\7\u201b\2\2\u01a0V\3\2\2\2\u01a1\u01a2\7\u00f1\2\2\u01a2"+
		"\u01a3\7\u00bb\2\2\u01a3\u01a4\7\u02de\2\2\u01a4X\3\2\2\2\u01a5\u01a6"+
		"\7\u00f1\2\2\u01a6\u01a7\7\u00bb\2\2\u01a7\u01a8\7\u00a5\2\2\u01a8Z\3"+
		"\2\2\2\u01a9\u01aa\7\u00f1\2\2\u01aa\u01ab\7\u00be\2\2\u01ab\u01ac\7\uffff"+
		"\2\2\u01ac\\\3\2\2\2\u01ad\u01ae\t\2\2\2\u01ae\u01af\t\3\2\2\u01af\u01b0"+
		"\t\4\2\2\u01b0\u01b1\t\5\2\2\u01b1\u01b2\t\3\2\2\u01b2^\3\2\2\2\u01b3"+
		"\u01b4\t\6\2\2\u01b4\u01b5\t\7\2\2\u01b5\u01b6\t\7\2\2\u01b6`\3\2\2\2"+
		"\u01b7\u01b8\t\5\2\2\u01b8\u01b9\t\b\2\2\u01b9\u01ba\t\t\2\2\u01ba\u01bb"+
		"\t\4\2\2\u01bb\u01bc\t\5\2\2\u01bc\u01bd\t\3\2\2\u01bd\u01be\t\6\2\2\u01be"+
		"\u01bf\t\7\2\2\u01bfb\3\2\2\2\u01c0\u01c1\t\n\2\2\u01c1\u01c2\t\6\2\2"+
		"\u01c2\u01c3\t\t\2\2\u01c3\u01c4\t\13\2\2\u01c4\u01c5\t\f\2\2\u01c5d\3"+
		"\2\2\2\u01c6\u01c7\t\2\2\2\u01c7\u01c8\t\3\2\2\u01c8\u01c9\t\r\2\2\u01c9"+
		"\u01ca\t\4\2\2\u01ca\u01cb\t\3\2\2\u01cb\u01cc\t\16\2\2\u01ccf\3\2\2\2"+
		"\u01cd\u01ce\t\6\2\2\u01ce\u01cf\t\17\2\2\u01cfh\3\2\2\2\u01d0\u01d1\t"+
		"\n\2\2\u01d1\u01d2\t\20\2\2\u01d2\u01d3\t\21\2\2\u01d3\u01d4\t\22\2\2"+
		"\u01d4\u01d5\t\20\2\2\u01d5j\3\2\2\2\u01d6\u01d7\t\5\2\2\u01d7\u01d8\t"+
		"\3\2\2\u01d8l\3\2\2\2\u01d9\u01da\t\13\2\2\u01da\u01db\t\21\2\2\u01db"+
		"\u01dc\t\20\2\2\u01dc\u01dd\t\6\2\2\u01dd\u01de\t\t\2\2\u01de\u01df\t"+
		"\20\2\2\u01dfn\3\2\2\2\u01e0\u01e1\t\17\2\2\u01e1\u01e2\t\20\2\2\u01e2"+
		"\u01e3\t\t\2\2\u01e3p\3\2\2\2\u01e4\u01e5\t\16\2\2\u01e5\u01e6\t\20\2"+
		"\2\u01e6\u01e7\t\t\2\2\u01e7\u01e8\t\6\2\2\u01e8\u01e9\t\13\2\2\u01e9"+
		"\u01ea\t\f\2\2\u01ear\3\2\2\2\u01eb\u01ec\t\16\2\2\u01ec\u01ed\t\20\2"+
		"\2\u01ed\u01ee\t\7\2\2\u01ee\u01ef\t\20\2\2\u01ef\u01f0\t\t\2\2\u01f0"+
		"\u01f1\t\20\2\2\u01f1t\3\2\2\2\u01f2\u01f3\t\21\2\2\u01f3\u01f4\t\20\2"+
		"\2\u01f4\u01f5\t\n\2\2\u01f5\u01f6\t\5\2\2\u01f6\u01f7\t\23\2\2\u01f7"+
		"\u01f8\t\20\2\2\u01f8v\3\2\2\2\u01f9\u01fa\t\13\2\2\u01fa\u01fb\t\6\2"+
		"\2\u01fb\u01fc\t\7\2\2\u01fc\u01fd\t\7\2\2\u01fdx\3\2\2\2\u01fe\u01ff"+
		"\t\24\2\2\u01ff\u0200\t\4\2\2\u0200\u0201\t\20\2\2\u0201\u0202\t\7\2\2"+
		"\u0202\u0203\t\16\2\2\u0203z\3\2\2\2\u0204\u0205\t\r\2\2\u0205\u0206\t"+
		"\4\2\2\u0206\u0207\t\t\2\2\u0207\u0208\t\f\2\2\u0208|\3\2\2\2\u0209\u020a"+
		"\t\16\2\2\u020a\u020b\t\4\2\2\u020b\u020c\t\17\2\2\u020c\u020d\t\t\2\2"+
		"\u020d\u020e\t\4\2\2\u020e\u020f\t\3\2\2\u020f\u0210\t\13\2\2\u0210\u0211"+
		"\t\t\2\2\u0211~\3\2\2\2\u0212\u0213\t\21\2\2\u0213\u0214\t\20\2\2\u0214"+
		"\u0215\t\t\2\2\u0215\u0216\t\2\2\2\u0216\u0217\t\21\2\2\u0217\u0218\t"+
		"\3\2\2\u0218\u0080\3\2\2\2\u0219\u021a\t\5\2\2\u021a\u021b\t\21\2\2\u021b"+
		"\u021c\t\16\2\2\u021c\u021d\t\20\2\2\u021d\u021e\t\21\2\2\u021e\u0082"+
		"\3\2\2\2\u021f\u0220\t\25\2\2\u0220\u0221\t\24\2\2\u0221\u0084\3\2\2\2"+
		"\u0222\u0223\t\17\2\2\u0223\u0224\t\26\2\2\u0224\u0225\t\4\2\2\u0225\u0226"+
		"\t\b\2\2\u0226\u0086\3\2\2\2\u0227\u0228\t\7\2\2\u0228\u0229\t\4\2\2\u0229"+
		"\u022a\t\n\2\2\u022a\u022b\t\4\2\2\u022b\u022c\t\t\2\2\u022c\u0088\3\2"+
		"\2\2\u022d\u022e\t\6\2\2\u022e\u022f\t\17\2\2\u022f\u0230\t\13\2\2\u0230"+
		"\u0231\t\20\2\2\u0231\u0232\t\3\2\2\u0232\u0233\t\16\2\2\u0233\u0234\t"+
		"\4\2\2\u0234\u0235\t\3\2\2\u0235\u0236\t\22\2\2\u0236\u008a\3\2\2\2\u0237"+
		"\u0238\t\6\2\2\u0238\u0239\t\17\2\2\u0239\u023a\t\13\2\2\u023a\u008c\3"+
		"\2\2\2\u023b\u023c\t\16\2\2\u023c\u023d\t\20\2\2\u023d\u023e\t\17\2\2"+
		"\u023e\u023f\t\13\2\2\u023f\u0240\t\20\2\2\u0240\u0241\t\3\2\2\u0241\u0242"+
		"\t\16\2\2\u0242\u0243\t\4\2\2\u0243\u0244\t\3\2\2\u0244\u0245\t\22\2\2"+
		"\u0245\u008e\3\2\2\2\u0246\u0247\t\16\2\2\u0247\u0248\t\20\2\2\u0248\u0249"+
		"\t\17\2\2\u0249\u024a\t\13\2\2\u024a\u0090\3\2\2\2\u024b\u024c\t\r\2\2"+
		"\u024c\u024d\t\f\2\2\u024d\u024e\t\20\2\2\u024e\u024f\t\21\2\2\u024f\u0250"+
		"\t\20\2\2\u0250\u0092\3\2\2\2\u0251\u0252\t\5\2\2\u0252\u0253\t\21\2\2"+
		"\u0253\u0094\3\2\2\2\u0254\u0255\t\27\2\2\u0255\u0256\t\5\2\2\u0256\u0257"+
		"\t\21\2\2\u0257\u0096\3\2\2\2\u0258\u0259\t\6\2\2\u0259\u025a\t\3\2\2"+
		"\u025a\u025b\t\16\2\2\u025b\u0098\3\2\2\2\u025c\u025d\t\3\2\2\u025d\u025e"+
		"\t\5\2\2\u025e\u025f\t\t\2\2\u025f\u009a\3\2\2\2\u0260\u0261\t\4\2\2\u0261"+
		"\u0262\t\3\2\2\u0262\u009c\3\2\2\2\u0263\u0264\t\17\2\2\u0264\u0265\t"+
		"\t\2\2\u0265\u0266\t\6\2\2\u0266\u0267\t\21\2\2\u0267\u0268\t\t\2\2\u0268"+
		"\u0269\t\17\2\2\u0269\u009e\3\2\2\2\u026a\u026b\t\20\2\2\u026b\u026c\t"+
		"\3\2\2\u026c\u026d\t\16\2\2\u026d\u026e\t\17\2\2\u026e\u00a0\3\2\2\2\u026f"+
		"\u0270\t\13\2\2\u0270\u0271\t\5\2\2\u0271\u0272\t\3\2\2\u0272\u0273\t"+
		"\t\2\2\u0273\u0274\t\6\2\2\u0274\u0275\t\4\2\2\u0275\u0276\t\3\2\2\u0276"+
		"\u0277\t\17\2\2\u0277\u00a2\3\2\2\2\u0278\u0279\t\4\2\2\u0279\u027a\t"+
		"\17\2\2\u027a\u00a4\3\2\2\2\u027b\u027c\t\3\2\2\u027c\u027d\t\2\2\2\u027d"+
		"\u027e\t\7\2\2\u027e\u027f\t\7\2\2\u027f\u00a6\3\2\2\2\u0280\u0281\t\13"+
		"\2\2\u0281\u0282\t\5\2\2\u0282\u0283\t\2\2\2\u0283\u0284\t\3\2\2\u0284"+
		"\u0285\t\t\2\2\u0285\u00a8\3\2\2\2\u0286\u0287\t\6\2\2\u0287\u0288\t\3"+
		"\2\2\u0288\u0289\t\24\2\2\u0289\u00aa\3\2\2\2\u028a\u028b\t\3\2\2\u028b"+
		"\u028c\t\5\2\2\u028c\u028d\t\3\2\2\u028d\u028e\t\20\2\2\u028e\u00ac\3"+
		"\2\2\2\u028f\u0290\t\17\2\2\u0290\u0291\t\4\2\2\u0291\u0292\t\3\2\2\u0292"+
		"\u0293\t\22\2\2\u0293\u0294\t\7\2\2\u0294\u0295\t\20\2\2\u0295\u00ae\3"+
		"\2\2\2\u0296\u0297\t\t\2\2\u0297\u0298\t\21\2\2\u0298\u0299\t\2\2\2\u0299"+
		"\u029a\t\20\2\2\u029a\u00b0\3\2\2\2\u029b\u029c\t\30\2\2\u029c\u029d\t"+
		"\6\2\2\u029d\u029e\t\7\2\2\u029e\u029f\t\17\2\2\u029f\u02a0\t\20\2\2\u02a0"+
		"\u00b2\3\2\2\2\u02a1\u02a2\t\20\2\2\u02a2\u02a3\t\27\2\2\u02a3\u02a4\t"+
		"\4\2\2\u02a4\u02a5\t\17\2\2\u02a5\u02a6\t\t\2\2\u02a6\u02a7\t\17\2\2\u02a7"+
		"\u00b4\3\2\2\2\u02a8\u02a9\t\13\2\2\u02a9\u02aa\t\6\2\2\u02aa\u02ab\t"+
		"\17\2\2\u02ab\u02ac\t\20\2\2\u02ac\u00b6\3\2\2\2\u02ad\u02ae\t\20\2\2"+
		"\u02ae\u02af\t\7\2\2\u02af\u02b0\t\17\2\2\u02b0\u02b1\t\20\2\2\u02b1\u00b8"+
		"\3\2\2\2\u02b2\u02b3\t\20\2\2\u02b3\u02b4\t\3\2\2\u02b4\u02b5\t\16\2\2"+
		"\u02b5\u00ba\3\2\2\2\u02b6\u02b7\t\r\2\2\u02b7\u02b8\t\f\2\2\u02b8\u02b9"+
		"\t\20\2\2\u02b9\u02ba\t\3\2\2\u02ba\u00bc\3\2\2\2\u02bb\u02bc\t\t\2\2"+
		"\u02bc\u02bd\t\f\2\2\u02bd\u02be\t\20\2\2\u02be\u02bf\t\3\2\2\u02bf\u00be"+
		"\3\2\2\2\u02c0\u02c5\7$\2\2\u02c1\u02c4\5\u011f\u0090\2\u02c2\u02c4\5"+
		"\u00c1a\2\u02c3\u02c1\3\2\2\2\u02c3\u02c2\3\2\2\2\u02c4\u02c7\3\2\2\2"+
		"\u02c5\u02c3\3\2\2\2\u02c5\u02c6\3\2\2\2\u02c6\u02c8\3\2\2\2\u02c7\u02c5"+
		"\3\2\2\2\u02c8\u02d3\7$\2\2\u02c9\u02ce\7)\2\2\u02ca\u02cd\5\u010b\u0086"+
		"\2\u02cb\u02cd\5\u00c1a\2\u02cc\u02ca\3\2\2\2\u02cc\u02cb\3\2\2\2\u02cd"+
		"\u02d0\3\2\2\2\u02ce\u02cc\3\2\2\2\u02ce\u02cf\3\2\2\2\u02cf\u02d1\3\2"+
		"\2\2\u02d0\u02ce\3\2\2\2\u02d1\u02d3\7)\2\2\u02d2\u02c0\3\2\2\2\u02d2"+
		"\u02c9\3\2\2\2\u02d3\u00c0\3\2\2\2\u02d4\u02e6\7^\2\2\u02d5\u02e7\t\31"+
		"\2\2\u02d6\u02d7\t\2\2\2\u02d7\u02d8\5\u00cbf\2\u02d8\u02d9\5\u00cbf\2"+
		"\u02d9\u02da\5\u00cbf\2\u02da\u02db\5\u00cbf\2\u02db\u02e7\3\2\2\2\u02dc"+
		"\u02dd\t\2\2\2\u02dd\u02de\5\u00cbf\2\u02de\u02df\5\u00cbf\2\u02df\u02e0"+
		"\5\u00cbf\2\u02e0\u02e1\5\u00cbf\2\u02e1\u02e2\5\u00cbf\2\u02e2\u02e3"+
		"\5\u00cbf\2\u02e3\u02e4\5\u00cbf\2\u02e4\u02e5\5\u00cbf\2\u02e5\u02e7"+
		"\3\2\2\2\u02e6\u02d5\3\2\2\2\u02e6\u02d6\3\2\2\2\u02e6\u02dc\3\2\2\2\u02e7"+
		"\u00c2\3\2\2\2\u02e8\u02e9\7\62\2\2\u02e9\u02ea\7z\2\2\u02ea\u02ec\3\2"+
		"\2\2\u02eb\u02ed\5\u00cbf\2\u02ec\u02eb\3\2\2\2\u02ed\u02ee\3\2\2\2\u02ee"+
		"\u02ec\3\2\2\2\u02ee\u02ef\3\2\2\2\u02ef\u00c4\3\2\2\2\u02f0\u02f9\5\u00d5"+
		"k\2\u02f1\u02f5\5\u00cfh\2\u02f2\u02f4\5\u00cdg\2\u02f3\u02f2\3\2\2\2"+
		"\u02f4\u02f7\3\2\2\2\u02f5\u02f3\3\2\2\2\u02f5\u02f6\3\2\2\2\u02f6\u02f9"+
		"\3\2\2\2\u02f7\u02f5\3\2\2\2\u02f8\u02f0\3\2\2\2\u02f8\u02f1\3\2\2\2\u02f9"+
		"\u00c6\3\2\2\2\u02fa\u02fc\5\u00d5k\2\u02fb\u02fd\5\u00d3j\2\u02fc\u02fb"+
		"\3\2\2\2\u02fd\u02fe\3\2\2\2\u02fe\u02fc\3\2\2\2\u02fe\u02ff\3\2\2\2\u02ff"+
		"\u00c8\3\2\2\2\u0300\u0302\t\32\2\2\u0301\u0300\3\2\2\2\u0302\u00ca\3"+
		"\2\2\2\u0303\u0306\5\u00cdg\2\u0304\u0306\5\u00c9e\2\u0305\u0303\3\2\2"+
		"\2\u0305\u0304\3\2\2\2\u0306\u00cc\3\2\2\2\u0307\u030a\5\u00d5k\2\u0308"+
		"\u030a\5\u00cfh\2\u0309\u0307\3\2\2\2\u0309\u0308\3\2\2\2\u030a\u00ce"+
		"\3\2\2\2\u030b\u030e\5\u00d1i\2\u030c\u030e\4:;\2\u030d\u030b\3\2\2\2"+
		"\u030d\u030c\3\2\2\2\u030e\u00d0\3\2\2\2\u030f\u0310\4\639\2\u0310\u00d2"+
		"\3\2\2\2\u0311\u0314\5\u00d5k\2\u0312\u0314\5\u00d1i\2\u0313\u0311\3\2"+
		"\2\2\u0313\u0312\3\2\2\2\u0314\u00d4\3\2\2\2\u0315\u0316\7\62\2\2\u0316"+
		"\u00d6\3\2\2\2\u0317\u0319\5\u00cdg\2\u0318\u0317\3\2\2\2\u0319\u031a"+
		"\3\2\2\2\u031a\u0318\3\2\2\2\u031a\u031b\3\2\2\2\u031b\u032e\3\2\2\2\u031c"+
		"\u031e\5\u00cdg\2\u031d\u031c\3\2\2\2\u031e\u031f\3\2\2\2\u031f\u031d"+
		"\3\2\2\2\u031f\u0320\3\2\2\2\u0320\u0321\3\2\2\2\u0321\u0323\7\60\2\2"+
		"\u0322\u0324\5\u00cdg\2\u0323\u0322\3\2\2\2\u0324\u0325\3\2\2\2\u0325"+
		"\u0323\3\2\2\2\u0325\u0326\3\2\2\2\u0326\u032e\3\2\2\2\u0327\u0329\7\60"+
		"\2\2\u0328\u032a\5\u00cdg\2\u0329\u0328\3\2\2\2\u032a\u032b\3\2\2\2\u032b"+
		"\u0329\3\2\2\2\u032b\u032c\3\2\2\2\u032c\u032e\3\2\2\2\u032d\u0318\3\2"+
		"\2\2\u032d\u031d\3\2\2\2\u032d\u0327\3\2\2\2\u032e\u0330\3\2\2\2\u032f"+
		"\u0331\t\20\2\2\u0330\u032f\3\2\2\2\u0331\u0333\3\2\2\2\u0332\u0334\7"+
		"/\2\2\u0333\u0332\3\2\2\2\u0333\u0334\3\2\2\2\u0334\u0336\3\2\2\2\u0335"+
		"\u0337\5\u00cdg\2\u0336\u0335\3\2\2\2\u0337\u0338\3\2\2\2\u0338\u0336"+
		"\3\2\2\2\u0338\u0339\3\2\2\2\u0339\u00d8\3\2\2\2\u033a\u033c\5\u00cdg"+
		"\2\u033b\u033a\3\2\2\2\u033c\u033f\3\2\2\2\u033d\u033b\3\2\2\2\u033d\u033e"+
		"\3\2\2\2\u033e\u0340\3\2\2\2\u033f\u033d\3\2\2\2\u0340\u0342\7\60\2\2"+
		"\u0341\u0343\5\u00cdg\2\u0342\u0341\3\2\2\2\u0343\u0344\3\2\2\2\u0344"+
		"\u0342\3\2\2\2\u0344\u0345\3\2\2\2\u0345\u00da\3\2\2\2\u0346\u0347\t\13"+
		"\2\2\u0347\u0348\t\5\2\2\u0348\u0349\t\3\2\2\u0349\u034a\t\17\2\2\u034a"+
		"\u034b\t\t\2\2\u034b\u034c\t\21\2\2\u034c\u034d\t\6\2\2\u034d\u034e\t"+
		"\4\2\2\u034e\u034f\t\3\2\2\u034f\u0350\t\t\2\2\u0350\u00dc\3\2\2\2\u0351"+
		"\u0352\t\16\2\2\u0352\u0353\t\5\2\2\u0353\u00de\3\2\2\2\u0354\u0355\t"+
		"\30\2\2\u0355\u0356\t\5\2\2\u0356\u0357\t\21\2\2\u0357\u00e0\3\2\2\2\u0358"+
		"\u0359\t\21\2\2\u0359\u035a\t\20\2\2\u035a\u035b\t\33\2\2\u035b\u035c"+
		"\t\2\2\2\u035c\u035d\t\4\2\2\u035d\u035e\t\21\2\2\u035e\u035f\t\20\2\2"+
		"\u035f\u00e2\3\2\2\2\u0360\u0361\t\2\2\2\u0361\u0362\t\3\2\2\u0362\u0363"+
		"\t\4\2\2\u0363\u0364\t\33\2\2\u0364\u0365\t\2\2\2\u0365\u0366\t\20\2\2"+
		"\u0366\u00e4\3\2\2\2\u0367\u0368\t\n\2\2\u0368\u0369\t\6\2\2\u0369\u036a"+
		"\t\3\2\2\u036a\u036b\t\16\2\2\u036b\u036c\t\6\2\2\u036c\u036d\t\t\2\2"+
		"\u036d\u036e\t\5\2\2\u036e\u036f\t\21\2\2\u036f\u0370\t\24\2\2\u0370\u00e6"+
		"\3\2\2\2\u0371\u0372\t\17\2\2\u0372\u0373\t\13\2\2\u0373\u0374\t\6\2\2"+
		"\u0374\u0375\t\7\2\2\u0375\u0376\t\6\2\2\u0376\u0377\t\21\2\2\u0377\u00e8"+
		"\3\2\2\2\u0378\u0379\t\5\2\2\u0379\u037a\t\30\2\2\u037a\u00ea\3\2\2\2"+
		"\u037b\u037c\t\6\2\2\u037c\u037d\t\16\2\2\u037d\u037e\t\16\2\2\u037e\u00ec"+
		"\3\2\2\2\u037f\u0380\t\16\2\2\u0380\u0381\t\21\2\2\u0381\u0382\t\5\2\2"+
		"\u0382\u0383\t\b\2\2\u0383\u00ee\3\2\2\2\u0384\u0385\t\30\2\2\u0385\u0386"+
		"\t\4\2\2\u0386\u0387\t\7\2\2\u0387\u0388\t\t\2\2\u0388\u0389\t\20\2\2"+
		"\u0389\u038a\t\21\2\2\u038a\u00f0\3\2\2\2\u038b\u038c\t\20\2\2\u038c\u038d"+
		"\t\27\2\2\u038d\u038e\t\t\2\2\u038e\u038f\t\21\2\2\u038f\u0390\t\6\2\2"+
		"\u0390\u0391\t\13\2\2\u0391\u0392\t\t\2\2\u0392\u00f2\3\2\2\2\u0393\u0397"+
		"\5\u00f5{\2\u0394\u0396\5\u00f7|\2\u0395\u0394\3\2\2\2\u0396\u0399\3\2"+
		"\2\2\u0397\u0395\3\2\2\2\u0397\u0398\3\2\2\2\u0398\u00f4\3\2\2\2\u0399"+
		"\u0397\3\2\2\2\u039a\u039d\5\u0127\u0094\2\u039b\u039d\5\u011b\u008e\2"+
		"\u039c\u039a\3\2\2\2\u039c\u039b\3\2\2\2\u039d\u00f6\3\2\2\2\u039e\u03a1"+
		"\5\u0107\u0084\2\u039f\u03a1\5\u0117\u008c\2\u03a0\u039e\3\2\2\2\u03a0"+
		"\u039f\3\2\2\2\u03a1\u00f8\3\2\2\2\u03a2\u03a6\7b\2\2\u03a3\u03a5\5\u0103"+
		"\u0082\2\u03a4\u03a3\3\2\2\2\u03a5\u03a8\3\2\2\2\u03a6\u03a4\3\2\2\2\u03a6"+
		"\u03a7\3\2\2\2\u03a7\u03a9\3\2\2\2\u03a8\u03a6\3\2\2\2\u03a9\u03ab\7b"+
		"\2\2\u03aa\u03a2\3\2\2\2\u03ab\u03ac\3\2\2\2\u03ac\u03aa\3\2\2\2\u03ac"+
		"\u03ad\3\2\2\2\u03ad\u00fa\3\2\2\2\u03ae\u03b0\5\u00fd\177\2\u03af\u03ae"+
		"\3\2\2\2\u03b0\u03b1\3\2\2\2\u03b1\u03af\3\2\2\2\u03b1\u03b2\3\2\2\2\u03b2"+
		"\u00fc\3\2\2\2\u03b3\u03f7\5\u0119\u008d\2\u03b4\u03f7\5\u011d\u008f\2"+
		"\u03b5\u03f7\5\u0121\u0091\2\u03b6\u03f7\5\u0123\u0092\2\u03b7\u03f7\5"+
		"\u0101\u0081\2\u03b8\u03f7\5\u0115\u008b\2\u03b9\u03f7\5\u0113\u008a\2"+
		"\u03ba\u03f7\5\u0111\u0089\2\u03bb\u03f7\5\u0105\u0083\2\u03bc\u03f7\5"+
		"\u0125\u0093\2\u03bd\u03be\7\u00e3\2\2\u03be\u03bf\7\u0163\2\2\u03bf\u03f7"+
		"\7\u20ae\2\2\u03c0\u03c1\7\u00e3\2\2\u03c1\u03c2\7\u00a2\2\2\u03c2\u03f7"+
		"\7\u017f\2\2\u03c3\u03c4\7\u00e4\2\2\u03c4\u03c5\7\u20ae\2\2\u03c5\u03f7"+
		"\7\u20ae\2\2\u03c6\u03c7\7\u00e4\2\2\u03c7\u03c8\7\u20ae\2\2\u03c8\u03f7"+
		"\7\uffff\2\2\u03c9\u03ca\7\u00e4\2\2\u03ca\u03cb\7\u20ae\2\2\u03cb\u03f7"+
		"\7\u201c\2\2\u03cc\u03cd\7\u00e4\2\2\u03cd\u03ce\7\u20ae\2\2\u03ce\u03f7"+
		"\7\u0194\2\2\u03cf\u03d0\7\u00e4\2\2\u03d0\u03d1\7\u20ae\2\2\u03d1\u03f7"+
		"\7\u2020\2\2\u03d2\u03d3\7\u00e4\2\2\u03d3\u03d4\7\u20ae\2\2\u03d4\u03f7"+
		"\7\u2028\2\2\u03d5\u03d6\7\u00e4\2\2\u03d6\u03d7\7\u20ae\2\2\u03d7\u03f7"+
		"\7\u2022\2\2\u03d8\u03d9\7\u00e4\2\2\u03d9\u03da\7\u20ae\2\2\u03da\u03f7"+
		"\7\u02c8\2\2\u03db\u03dc\7\u00e4\2\2\u03dc\u03dd\7\u20ae\2\2\u03dd\u03f7"+
		"\7\u2032\2\2\u03de\u03df\7\u00e4\2\2\u03df\u03e0\7\u20ae\2\2\u03e0\u03f7"+
		"\7\u0162\2\2\u03e1\u03e2\7\u00e4\2\2\u03e2\u03e3\7\u20ae\2\2\u03e3\u03f7"+
		"\7\u00aa\2\2\u03e4\u03e5\7\u00e4\2\2\u03e5\u03e6\7\u20ae\2\2\u03e6\u03f7"+
		"\7\u00ab\2\2\u03e7\u03e8\7\u00e4\2\2\u03e8\u03e9\7\uffff\2\2\u03e9\u03f7"+
		"\7\u017a\2\2\u03ea\u03eb\7\u00e5\2\2\u03eb\u03ec\7\u20ae\2\2\u03ec\u03f7"+
		"\7\u20ae\2\2\u03ed\u03ee\7\u00c4\2\2\u03ee\u03f7\7\u00a2\2\2\u03ef\u03f0"+
		"\7\u00e4\2\2\u03f0\u03f1\7\u20ae\2\2\u03f1\u03f7\7\u2023\2\2\u03f2\u03f3"+
		"\7\u00e4\2\2\u03f3\u03f4\7\u20ae\2\2\u03f4\u03f7\7\u00b1\2\2\u03f5\u03f7"+
		"\5\u00ff\u0080\2\u03f6\u03b3\3\2\2\2\u03f6\u03b4\3\2\2\2\u03f6\u03b5\3"+
		"\2\2\2\u03f6\u03b6\3\2\2\2\u03f6\u03b7\3\2\2\2\u03f6\u03b8\3\2\2\2\u03f6"+
		"\u03b9\3\2\2\2\u03f6\u03ba\3\2\2\2\u03f6\u03bb\3\2\2\2\u03f6\u03bc\3\2"+
		"\2\2\u03f6\u03bd\3\2\2\2\u03f6\u03c0\3\2\2\2\u03f6\u03c3\3\2\2\2\u03f6"+
		"\u03c6\3\2\2\2\u03f6\u03c9\3\2\2\2\u03f6\u03cc\3\2\2\2\u03f6\u03cf\3\2"+
		"\2\2\u03f6\u03d2\3\2\2\2\u03f6\u03d5\3\2\2\2\u03f6\u03d8\3\2\2\2\u03f6"+
		"\u03db\3\2\2\2\u03f6\u03de\3\2\2\2\u03f6\u03e1\3\2\2\2\u03f6\u03e4\3\2"+
		"\2\2\u03f6\u03e7\3\2\2\2\u03f6\u03ea\3\2\2\2\u03f6\u03ed\3\2\2\2\u03f6"+
		"\u03ef\3\2\2\2\u03f6\u03f2\3\2\2\2\u03f6\u03f5\3\2\2\2\u03f7\u00fe\3\2"+
		"\2\2\u03f8\u03f9\7\61\2\2\u03f9\u03fa\7,\2\2\u03fa\u0400\3\2\2\2\u03fb"+
		"\u03ff\5\u0109\u0085\2\u03fc\u03fd\7,\2\2\u03fd\u03ff\5\u010f\u0088\2"+
		"\u03fe\u03fb\3\2\2\2\u03fe\u03fc\3\2\2\2\u03ff\u0402\3\2\2\2\u0400\u03fe"+
		"\3\2\2\2\u0400\u0401\3\2\2\2\u0401\u0403\3\2\2\2\u0402\u0400\3\2\2\2\u0403"+
		"\u0404\7,\2\2\u0404\u0416\7\61\2\2\u0405\u0406\7\61\2\2\u0406\u0407\7"+
		"\61\2\2\u0407\u040b\3\2\2\2\u0408\u040a\5\u010d\u0087\2\u0409\u0408\3"+
		"\2\2\2\u040a\u040d\3\2\2\2\u040b\u0409\3\2\2\2\u040b\u040c\3\2\2\2\u040c"+
		"\u040f\3\2\2\2\u040d\u040b\3\2\2\2\u040e\u0410\5\u0115\u008b\2\u040f\u040e"+
		"\3\2\2\2\u040f\u0410\3\2\2\2\u0410\u0413\3\2\2\2\u0411\u0414\5\u0121\u0091"+
		"\2\u0412\u0414\7\2\2\3\u0413\u0411\3\2\2\2\u0413\u0412\3\2\2\2\u0414\u0416"+
		"\3\2\2\2\u0415\u03f8\3\2\2\2\u0415\u0405\3\2\2\2\u0416\u0100\3\2\2\2\u0417"+
		"\u0418\t\34\2\2\u0418\u0102\3\2\2\2\u0419\u041a\t\35\2\2\u041a\u0104\3"+
		"\2\2\2\u041b\u041c\t\36\2\2\u041c\u0106\3\2\2\2\u041d\u041e\t\37\2\2\u041e"+
		"\u0108\3\2\2\2\u041f\u0420\t \2\2\u0420\u010a\3\2\2\2\u0421\u0422\t!\2"+
		"\2\u0422\u010c\3\2\2\2\u0423\u0424\t\"\2\2\u0424\u010e\3\2\2\2\u0425\u0426"+
		"\t#\2\2\u0426\u0110\3\2\2\2\u0427\u0428\t$\2\2\u0428\u0112\3\2\2\2\u0429"+
		"\u042a\t%\2\2\u042a\u0114\3\2\2\2\u042b\u042c\t&\2\2\u042c\u0116\3\2\2"+
		"\2\u042d\u042e\t\'\2\2\u042e\u0118\3\2\2\2\u042f\u0430\t(\2\2\u0430\u011a"+
		"\3\2\2\2\u0431\u0432\t)\2\2\u0432\u011c\3\2\2\2\u0433\u0434\t*\2\2\u0434"+
		"\u011e\3\2\2\2\u0435\u0436\t+\2\2\u0436\u0120\3\2\2\2\u0437\u0438\t,\2"+
		"\2\u0438\u0122\3\2\2\2\u0439\u043a\t-\2\2\u043a\u0124\3\2\2\2\u043b\u043c"+
		"\t.\2\2\u043c\u0126\3\2\2\2\u043d\u043e\t/\2\2\u043e\u0128\3\2\2\2)\2"+
		"\u02c3\u02c5\u02cc\u02ce\u02d2\u02e6\u02ee\u02f5\u02f8\u02fe\u0301\u0305"+
		"\u0309\u030d\u0313\u031a\u031f\u0325\u032b\u032d\u0330\u0333\u0338\u033d"+
		"\u0344\u0397\u039c\u03a0\u03a6\u03ac\u03b1\u03f6\u03fe\u0400\u040b\u040f"+
		"\u0413\u0415\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}